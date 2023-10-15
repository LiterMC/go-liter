
package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
)

type Conn struct {
	*liter.Conn
	When time.Time
	player atomic.Pointer[liter.PlayerInfo]
}

func (c *Conn)ActiveTime()(time.Duration){
	return time.Since(c.When)
}

func (c *Conn)SetPlayer(player liter.PlayerInfo){
	c.player.Store(&player)
}

func (c *Conn)Player()(*liter.PlayerInfo){
	return c.player.Load()
}

type Server struct{
	Addr string
	scripts *script.Manager
	handler http.Handler
	users   *UserStorage

	inShutdown atomic.Bool
	mux        sync.Mutex
	listeners  []net.Listener
	conns      FlatMemory[*Conn]
}

func NewServer(sm *script.Manager)(s *Server){
	s = &Server{
		scripts: sm,
		users: NewUserStorage(filepath.Join(configDir, "users.json")),
	}
	s.initHandler()
	if err := s.users.Load(); errors.Is(err, os.ErrNotExist) {
		s.users.Save()
	}
	return
}

func (s *Server)Scripts()(*script.Manager){
	return s.scripts
}

func (s *Server)closeListenersLocked()(err error){
	for _, l := range s.listeners {
		if e := l.Close(); e != nil {
			if err == nil {
				err = e
			}
		}
	}
	s.listeners = nil
	return
}

func (s *Server)Serve(listener net.Listener)(err error){
	s.mux.Lock()
	s.listeners = append(s.listeners, listener)
	s.mux.Unlock()

	for {
		var c net.Conn
		if c, err = listener.Accept(); err != nil {
			if errors.Is(err, net.ErrClosed) {
				err = nil
			}
			return
		}
		go func(c net.Conn){
			defer func(){
				if err := recover(); err != nil {
					loger.Errorf("Error while handling %s:\n%v\n%s", c.RemoteAddr().String(), err, getStacktrace())
				}
			}()

			successed := false
			defer func(){
				if !successed {
					c.Close()
				}
			}()
			cfg, _ := getConfig()
			if host, _, err := net.SplitHostPort(c.RemoteAddr().String()); err != nil {
				return
			}else{
				ip := net.ParseIP(host)
				if ip == nil {
					return
				}
				if blacklist.IncludeIP(ip) {
					return
				}
				if cfg.EnableIPWhitelist {
					if !whitelist.IncludeIP(ip) {
						return
					}
				}
			}
			wc := liter.WrapConn(c)
			successed = true
			s.handle(wc, &cfg)
		}(c)
	}
}

func (s *Server)Shutdown(ctx context.Context)(err error){
	s.inShutdown.Store(true)

	s.mux.Lock()
	s.closeListenersLocked()
	s.mux.Unlock()

	if s.conns.Count() == 0 {
		return
	}
	select {
	case <-time.After(time.Millisecond * 50):
		if s.conns.Count() == 0 {
			return
		}
	case <-ctx.Done():
		err = ctx.Err()
		s.mux.Lock()
		defer s.mux.Unlock()
		if s.conns.Count() > 0 {
			s.conns.ForEach(func(_ int, c *Conn){
				c.Close()
			})
			s.conns.Clear()
		}
	}
	return
}
