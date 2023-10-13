
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenIdSet = make(map[string]struct{})
	tokenIdSetMux sync.RWMutex
)

func registerTokenId(id string){
	tokenIdSetMux.Lock()
	defer tokenIdSetMux.Unlock()
	tokenIdSet[id] = struct{}{}
}

func unregisterTokenId(id string){
	tokenIdSetMux.Lock()
	defer tokenIdSetMux.Unlock()
	delete(tokenIdSet, id)
}

func checkTokenId(id string)(ok bool){
	tokenIdSetMux.RLock()
	defer tokenIdSetMux.RUnlock()
	_, ok = tokenIdSet[id]
	return
}

func (s *Server)checkToken(ctx *gin.Context, token string)(ok bool){
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error){
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return hmacKey, nil
		},
		jwt.WithSubject("optk"),
		jwt.WithIssuedAt(),
		jwt.WithIssuer(jwtIssuer),
	)
	if err != nil {
		loger.Debugf("JWT verify error: %v", err)
		return false
	}
	c, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		loger.Debugf("JWT claim is not jwt.MapClaims")
		return false
	}
	id := ctx.GetString(clientIdKey)
	if c["cli"] != id {
		loger.Debugf("JWT cli %q not match %q", c["cli"], id)
		return false
	}
	if jti, ok := c["jti"].(string); !ok || !checkTokenId(jti) {
		ctx.Set(clientTokenIdKey, jti)
		loger.Debug("JWT id not exist")
		return false
	}
	if u, ok := c["user"]; !ok {
		return false
	}else{
		ctx.Set(clientUserKey, u)
	}
	return true
}

// checkTokenMiddle method will verify the token defined in the http header "X-Token"
// use custom header instead of standard "Authorization" is for avoid XSS attacks
func (s *Server)checkTokenMiddle(ctx *gin.Context){
	token := ctx.GetHeader("X-Token")
	if !s.checkToken(ctx, token) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, RequestFailedFromString(
			"AuthError", "token expired",
		))
		return
	}
	ctx.Next()
}

func (s *Server)initV1(v1 *gin.RouterGroup){
	v1.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1.POST("/login", func(ctx *gin.Context){
		var req struct{
			User   string `json:"username"`
			Passwd string `json:"password"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}

		u := s.users.GetUser(req.User)
		if u == nil || !u.CheckPassword(req.Passwd) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, RequestFailedFromString(
				"AuthError", "Username or password is error",
			))
			return
		}

		id := ctx.GetString(clientIdKey)
		jti, err := genRandB64(16)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		now := time.Now()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"jti": jti,
			"sub": "optk",
			"iss": jwtIssuer,
			"iat": toSeconds(now),
			"exp": toSeconds(now.Add(time.Hour * 12)),
			"cli": id,
			"user": req.User,
		})
		tokenStr, err := token.SignedString(hmacKey)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		registerTokenId(jti)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"token": tokenStr,
		})
	})

	v1.POST("/logout", s.checkTokenMiddle, func(ctx *gin.Context){
		unregisterTokenId(ctx.GetString(clientTokenIdKey))
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1.GET("/verify", s.checkTokenMiddle, func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1.POST("/changepasswd", s.checkTokenMiddle, func(ctx *gin.Context){
		user := ctx.GetString(clientUserKey)

		var req struct {
			OldPasswd string `json:"oldPassword"`
			NewPasswd string `json:"newPassword"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}

		u := s.users.GetUser(user)
		ok := u.CheckPassword(req.OldPasswd)
		if u == nil || !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, RequestFailedFromString(
				"AuthError", "Password is error",
			))
			return
		}
		u.SetPassword(req.NewPasswd)
		s.users.Save()

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	registerConfigs(s, v1)
}

func registerConfigs(s *Server, g *gin.RouterGroup){
	g.GET("/config", s.checkTokenMiddle, func(ctx *gin.Context){
		cfg, cfgHash := getConfig()
		ctx.Header("ETag", strconv.Quote(cfgHash))
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == cfgHash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"onlineMode": cfg.OnlineMode,
			"enableWhitelist": cfg.EnableWhitelist,
			"enableIPWhitelist": cfg.EnableIPWhitelist,
		})
	})

	g.POST("/config", s.checkTokenMiddle, func(ctx *gin.Context){
		configLock.Lock()
		defer configLock.Unlock()

		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		}else if savedHash != strconv.Quote(cfgHash) {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		var req struct {
			Op string `json:"op"`
			Value json.RawMessage `json:"value"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}

		var tmp Config // we define a temporay config variable here to avoid bad data break the global config
		switch req.Op {
		case "onlineMode":
			if err := json.Unmarshal(req.Value, &tmp.OnlineMode); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
				return
			}
			config.OnlineMode = tmp.OnlineMode
		case "enableWhitelist":
			if err := json.Unmarshal(req.Value, &tmp.EnableWhitelist); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
				return
			}
			config.EnableWhitelist = tmp.EnableWhitelist
		case "enableIPWhitelist":
			if err := json.Unmarshal(req.Value, &tmp.EnableIPWhitelist); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
				return
			}
			config.EnableIPWhitelist = tmp.EnableIPWhitelist
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported config key %q", req.Op),
			))
			return
		}

		cfgHash, _ = genRandB64(48)
		if err := config.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	g.GET("/whitelist", s.checkTokenMiddle, func(ctx *gin.Context){
		wl, listHash := getWhitelist()
		ctx.Header("ETag", strconv.Quote(listHash))
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == listHash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": wl,
		})
	})

	g.POST("/whitelist", s.checkTokenMiddle, func(ctx *gin.Context){
		configLock.Lock()
		defer configLock.Unlock()

		listHash := cfgHash
		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		}else if savedHash != strconv.Quote(listHash) {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		const (
			opAddPlayer = "addpl"
			opRemovePlayer = "rmpl"
			opAddIP = "addip"
			opRemoveIP = "rmip"
		)

		var req struct {
			Op string `json:"op"`
			Value string `json:"value,omitempty"` // use when add
			Index int    `json:"index,omitempty"` // use when remove
		}
		req.Index = -1
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}

		switch req.Op {
		case opAddPlayer:
			if err := whitelist.AddPlayer(req.Value); err != nil {
				ctx.AbortWithError(http.StatusOK, err)
				return
			}
			whitelist.cleanPlayers()
		case opRemovePlayer:
			if req.Index < 0 || req.Index >= len(whitelist.Players) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			whitelist.Players = remove(whitelist.Players, req.Index)
		case opAddIP:
			if err := whitelist.AddIP(req.Value); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
				return
			}
		case opRemoveIP:
			if req.Index < 0 || req.Index >= len(whitelist.IPs) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			whitelist.IPs = remove(whitelist.IPs, req.Index)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported operation %q", req.Op),
			))
			return
		}

		cfgHash, _ = genRandB64(48)
		if err := whitelist.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	g.GET("/blacklist", s.checkTokenMiddle, func(ctx *gin.Context){
		wl, listHash := getBlacklist()
		ctx.Header("ETag", strconv.Quote(listHash))
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == listHash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": wl,
		})
	})

	g.POST("/blacklist", s.checkTokenMiddle, func(ctx *gin.Context){
		configLock.Lock()
		defer configLock.Unlock()

		listHash := cfgHash
		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		}else if savedHash != strconv.Quote(listHash) {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		const (
			opAddPlayer = "addpl"
			opRemovePlayer = "rmpl"
			opAddIP = "addip"
			opRemoveIP = "rmip"
		)

		var req struct {
			Op string `json:"op"`
			Value string `json:"value,omitempty"` // use when add
			Index int    `json:"index,omitempty"` // use when remove
		}
		req.Index = -1
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}

		switch req.Op {
		case opAddPlayer:
			if err := blacklist.AddPlayer(req.Value); err != nil {
				ctx.AbortWithError(http.StatusOK, err)
				return
			}
			blacklist.cleanPlayers()
		case opRemovePlayer:
			if req.Index < 0 || req.Index >= len(blacklist.Players) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			blacklist.Players = remove(blacklist.Players, req.Index)
		case opAddIP:
			if err := blacklist.AddIP(req.Value); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("EncodeError"))
				return
			}
		case opRemoveIP:
			if req.Index < 0 || req.Index >= len(blacklist.IPs) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			blacklist.IPs = remove(blacklist.IPs, req.Index)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported operation %q", req.Op),
			))
			return
		}

		cfgHash, _ = genRandB64(48)
		if err := blacklist.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})	
}