package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kmcsr/go-liter"
)

var (
	tokenIdSet    = make(map[string]struct{})
	tokenIdSetMux sync.RWMutex
)

func registerTokenId(id string) {
	tokenIdSetMux.Lock()
	defer tokenIdSetMux.Unlock()
	tokenIdSet[id] = struct{}{}
}

func unregisterTokenId(id string) {
	tokenIdSetMux.Lock()
	defer tokenIdSetMux.Unlock()
	delete(tokenIdSet, id)
}

func checkTokenId(id string) (ok bool) {
	tokenIdSetMux.RLock()
	defer tokenIdSetMux.RUnlock()
	_, ok = tokenIdSet[id]
	return
}

func (s *Server) checkToken(ctx *gin.Context, token string) (ok bool) {
	t, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return s.hmacKey, nil
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
	} else {
		ctx.Set(clientUserKey, u)
	}
	return true
}

// checkTokenMiddle method will verify the token defined in the http header "X-Token"
// use custom header instead of standard "Authorization" is for avoid XSS attacks
func (s *Server) checkTokenMiddle(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if !s.checkToken(ctx, token) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, RequestFailedFromString(
			"AuthError", "token expired",
		))
		return
	}
	ctx.Next()
}

func (s *Server) initV1(v1 *gin.RouterGroup) {
	checkedG := v1.Group("/", s.checkTokenMiddle)

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1.POST("/login", func(ctx *gin.Context) {
		var req struct {
			User   string `json:"username"`
			Passwd string `json:"password"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
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
			"jti":  jti,
			"sub":  "optk",
			"iss":  jwtIssuer,
			"iat":  toSeconds(now),
			"exp":  toSeconds(now.Add(time.Hour * 12)),
			"cli":  id,
			"user": req.User,
		})
		tokenStr, err := token.SignedString(s.hmacKey)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		registerTokenId(jti)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"token":  tokenStr,
		})
	})

	checkedG.POST("/logout", func(ctx *gin.Context) {
		unregisterTokenId(ctx.GetString(clientTokenIdKey))
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	checkedG.GET("/verify", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	checkedG.POST("/changepasswd", func(ctx *gin.Context) {
		user := ctx.GetString(clientUserKey)

		var req struct {
			OldPasswd string `json:"oldPassword"`
			NewPasswd string `json:"newPassword"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
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

	registerPlayerAPI(s, v1)
	registerConfigs(s, checkedG)
	registerStatus(s, checkedG)
}

// config APIs
func registerConfigs(s *Server, g *gin.RouterGroup) {
	g.GET("/config", func(ctx *gin.Context) {
		qhash := strconv.Quote(*s.cfgHash)
		ctx.Header("ETag", qhash)
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == qhash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":            "ok",
			"onlineMode":        s.config.OnlineMode,
			"enableWhitelist":   s.config.EnableWhitelist,
			"enableIPWhitelist": s.config.EnableIPWhitelist,
		})
	})

	g.POST("/config", func(ctx *gin.Context) {
		s.configLock.Lock()
		defer s.configLock.Unlock()

		qhash := strconv.Quote(*s.cfgHash)

		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		} else if savedHash != qhash {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		var req struct {
			Op    string          `json:"op"`
			Value json.RawMessage `json:"value"`
		}

		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
			return
		}

		var tmp Config // we define a temporay config variable here to avoid bad data break the global config
		switch req.Op {
		case "onlineMode":
			if err := json.Unmarshal(req.Value, &tmp.OnlineMode); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
				return
			}
			s.config.OnlineMode = tmp.OnlineMode
		case "enableWhitelist":
			if err := json.Unmarshal(req.Value, &tmp.EnableWhitelist); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
				return
			}
			s.config.EnableWhitelist = tmp.EnableWhitelist
		case "enableIPWhitelist":
			if err := json.Unmarshal(req.Value, &tmp.EnableIPWhitelist); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
				return
			}
			s.config.EnableIPWhitelist = tmp.EnableIPWhitelist
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported config key %q", req.Op),
			))
			return
		}

		*s.cfgHash, _ = genRandB64(48)
		if err := s.config.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	g.GET("/whitelist", func(ctx *gin.Context) {
		listHash := strconv.Quote(*s.cfgHash)
		ctx.Header("ETag", listHash)
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == listHash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   s.whitelist,
		})
	})

	g.POST("/whitelist", func(ctx *gin.Context) {
		s.configLock.Lock()
		defer s.configLock.Unlock()

		listHash := strconv.Quote(*s.cfgHash)
		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		} else if savedHash != listHash {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		const (
			opAddPlayer    = "addpl"
			opRemovePlayer = "rmpl"
			opAddIP        = "addip"
			opRemoveIP     = "rmip"
		)

		var req struct {
			Op    string `json:"op"`
			Value string `json:"value,omitempty"` // use when add
			Index int    `json:"index,omitempty"` // use when remove
		}
		req.Index = -1
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
			return
		}

		switch req.Op {
		case opAddPlayer:
			if err := s.whitelist.AddPlayer(req.Value, s.config.OnlineMode); err != nil {
				ctx.AbortWithError(http.StatusOK, err)
				return
			}
			s.whitelist.cleanPlayers()
		case opRemovePlayer:
			if req.Index < 0 || req.Index >= len(s.whitelist.Players) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			s.whitelist.Players = remove(s.whitelist.Players, req.Index)
		case opAddIP:
			if err := s.whitelist.AddIP(req.Value); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
				return
			}
		case opRemoveIP:
			if req.Index < 0 || req.Index >= len(s.whitelist.IPs) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			s.whitelist.IPs = remove(s.whitelist.IPs, req.Index)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported operation %q", req.Op),
			))
			return
		}

		*s.cfgHash, _ = genRandB64(48)
		if err := s.whitelist.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	g.GET("/blacklist", func(ctx *gin.Context) {
		listHash := strconv.Quote(*s.cfgHash)
		ctx.Header("ETag", listHash)
		if savedHash := ctx.GetHeader("If-None-Match"); len(savedHash) != 0 && savedHash == listHash {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   s.blacklist,
		})
	})

	g.POST("/blacklist", func(ctx *gin.Context) {
		s.configLock.Lock()
		defer s.configLock.Unlock()

		listHash := strconv.Quote(*s.cfgHash)
		if savedHash := ctx.GetHeader("If-Match"); len(savedHash) == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"HeaderMissing",
				`"If-Match" header required`,
			))
			return
		} else if savedHash != listHash {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, RequestFailedFromString(
				"ResourceModified", "",
			))
			return
		}

		const (
			opAddPlayer    = "addpl"
			opRemovePlayer = "rmpl"
			opAddIP        = "addip"
			opRemoveIP     = "rmip"
		)

		var req struct {
			Op    string `json:"op"`
			Value string `json:"value,omitempty"` // use when add
			Index int    `json:"index,omitempty"` // use when remove
		}
		req.Index = -1
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
			return
		}

		switch req.Op {
		case opAddPlayer:
			if err := s.blacklist.AddPlayer(req.Value, s.config.OnlineMode); err != nil {
				ctx.AbortWithError(http.StatusOK, err)
				return
			}
			s.blacklist.cleanPlayers()
		case opRemovePlayer:
			if req.Index < 0 || req.Index >= len(s.blacklist.Players) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			s.blacklist.Players = remove(s.blacklist.Players, req.Index)
		case opAddIP:
			if err := s.blacklist.AddIP(req.Value); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
				return
			}
		case opRemoveIP:
			if req.Index < 0 || req.Index >= len(s.blacklist.IPs) {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromString(
					"IndexOutOfBounds", fmt.Sprintf("Array index %d out of bounds", req.Index),
				))
				return
			}
			s.blacklist.IPs = remove(s.blacklist.IPs, req.Index)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromString(
				"UnknownOperation", fmt.Sprintf("Unsupported operation %q", req.Op),
			))
			return
		}

		*s.cfgHash, _ = genRandB64(48)
		if err := s.blacklist.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("SaveError"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

// register player api, for get basic informations
func registerPlayerAPI(s *Server, g *gin.RouterGroup) {
	const uuidKey = "liter.param.uuid"
	gu := g.Group("/player/:uuid/")
	gu.Use(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, RequestFailedFromError(err).SetType("DecodeError"))
			return
		}
		ctx.Set(uuidKey, id)
		ctx.Next()
	})
	gu.GET("/head", func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache")
		uid := ctx.Value(uuidKey).(uuid.UUID)
		profile, err := AuthClient.GetPlayerProfile(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("RemoteError"))
			return
		}
		skin, err := profile.GetSkin(AuthClient.Client)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("RemoteError"))
			return
		}
		ctx.Header("Etag", skin.Etag)
		if inm := ctx.GetHeader("If-None-Match"); len(inm) > 0 && inm == skin.Etag {
			ctx.Status(http.StatusNotModified)
			return
		}
		var buf bytes.Buffer
		if err = png.Encode(&buf, skin.Head()); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, RequestFailedFromError(err).SetType("EncodeError"))
			return
		}
		ctx.Data(http.StatusOK, "image/png", buf.Bytes())
	})
}

// server status APIs
func registerStatus(s *Server, g *gin.RouterGroup) {
	g.GET("/conns", func(ctx *gin.Context) {
		type resT struct {
			Id        int               `json:"id"`
			Addr      string            `json:"addr"`
			When      int64             `json:"when"`
			LocalAddr string            `json:"localAddr"`
			Server    string            `json:"server"`
			Player    *liter.PlayerInfo `json:"player,omitempty"`
		}
		ctx.Header("Cache-Control", "no-cache")
		lm := s.conns.LastModified().Format(time.RFC1123)
		data := make([]resT, 0, s.conns.Count())
		s.conns.ForEach(func(i int, conn *Conn) {
			local, svr, ok := conn.LocalServer()
			if !ok {
				return
			}
			data = append(data, resT{
				Id:        i,
				Addr:      conn.RemoteAddr().String(),
				When:      conn.When.Unix(),
				Player:    conn.Player(),
				LocalAddr: local,
				Server:    svr,
			})
		})
		ctx.Header("Last-Modified", lm)
		if lm2 := ctx.GetHeader("Last-Modified"); len(lm2) != 0 && lm2 == lm {
			ctx.Status(http.StatusNotModified)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   data,
		})
	})
}
