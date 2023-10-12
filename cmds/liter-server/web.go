
package main

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenExpired = errors.New("Operation token is expired")
)

func init(){
	gin.DefaultWriter = os.Stdout
}

var _ http.Handler = (*Server)(nil)

func (s *Server)ServeHTTP(rw http.ResponseWriter, req *http.Request){
	s.handler.ServeHTTP(rw, req)
}

func (s *Server)initHandler(){
	app := gin.Default()
	app.StaticFS("/main", DashboardAssets)
	s.initAPI(app.Group("/api"))
	app.GET("/", func(ctx *gin.Context){
		ctx.Redirect(http.StatusFound, "/main")
	})
	s.handler = app.Handler()
}

const (
	jwtIssuer = "litermc-webapi"
)

const (
	clientIdKey = "liter.client.id"
	clientUserKey = "liter.client.user"
)

var hmacKey []byte = func()(key []byte){
	path := filepath.Join(configDir, "server.hmac.private_key")
	buf, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			var sbuf string
			if sbuf, err = genRandB64(256); err != nil {
				loger.Fatalf("Cannot generate hmac key: %v", err)
			}
			buf = ([]byte)(sbuf)
			if err = os.WriteFile(path, buf, 0600); err != nil {
				loger.Fatalf("Cannot create hmac key file: %v", err)
			}
		}else{
			loger.Fatalf("Cannot read hmac key: %v", err)
		}
	}
	key = make([]byte, base64.RawURLEncoding.DecodedLen(len(buf)))
	if _, err = base64.RawURLEncoding.Decode(key, buf); err != nil {
		loger.Fatalf("Cannot decode hmac key: %v", err)
	}
	return
}()

func (s *Server)initAPI(api *gin.RouterGroup){
	api.Use(func(ctx *gin.Context){
		ctx.Next()
		if ctx.IsAborted() {
			if e := ctx.Errors.Last(); e != nil {
				ctx.JSON(http.StatusInternalServerError, RequestFailedFromError(e.Err))
			}
		}
	})
	api.Use(func(ctx *gin.Context){
		const cliIdCookieName = "_cliId"
		const cliIdLeng = 64
		var id string
		if cliId, err := ctx.Cookie(cliIdCookieName); err == nil {
			if base64.RawURLEncoding.DecodedLen(len(cliId)) == cliIdLeng {
				id = cliId
			}
		}
		if len(id) == 0 {
			var err error
			if id, err = genRandB64(cliIdLeng); err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
		ctx.SetCookie(cliIdCookieName, id, 60 * 60 * 24 * 30 * 356, "/", "", true, true) // a year
		ctx.Set(clientIdKey, asSha256(id)) // ensure the ID is secret
		ctx.Next()
	})
	s.initV1(api.Group("v1"))
}
