
package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
)

type Server struct{
	Addr string
	scripts *script.Manager
	handler http.Handler

	inShutdown atomic.Bool
	mux        sync.Mutex
	cond       *sync.Cond
	listeners  []net.Listener
	conns      []*liter.Conn
}

func NewServer(sm *script.Manager)(s *Server){
	s = &Server{
		scripts: sm,
	}
	s.cond = sync.NewCond(&s.mux)
	s.initHandler()
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
			successed := false
			defer func(){
				if !successed {
					c.Close()
				}
			}()
			cfg := getConfig()
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
			s.mux.Lock()
			s.conns = append(s.conns, wc)
			s.mux.Unlock()
			defer func(){
				s.mux.Lock()
				defer s.mux.Unlock()
				e := len(s.conns) - 1
				for i, d := range s.conns {
					if d == wc {
						if i != e {
							s.conns[i], s.conns[e] = s.conns[e], s.conns[i]
						}
						s.conns = s.conns[:e]
						if len(s.conns) == 0 {
							s.cond.Broadcast()
						}
						return
					}
				}
			}()
			successed = true
			s.handle(wc)
		}(c)
	}
}

func (s *Server)Shutdown(ctx context.Context)(err error){
	s.inShutdown.Store(true)

	s.mux.Lock()
	s.closeListenersLocked()
	s.mux.Unlock()

	select {
	case <-waitUntilNot(s.cond, func()(bool){
		return len(s.conns) > 0
	}):
	case <-ctx.Done():
		err = ctx.Err()
		s.mux.Lock()
		defer s.mux.Unlock()
		if len(s.conns) > 0 {
			for _, c := range s.conns {
				c.Close()
			}
			s.conns = nil
			s.cond.Broadcast()
		}
	}
	return
}
