
package main

import (
	"context"
	"errors"
	"net"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/net/proxy"
	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
	// "github.com/kmcsr/go-liter/packets"
)

var manager = script.NewManager()

func main(){
	scriptpath := filepath.Join(configDir, "plugins")
	manager.SetLogger(loger)
	var err error

RESTART:
	if _, err := manager.LoadFromDir(scriptpath); err != nil {
		loger.Errorf("Cannot load scripts: %v", err)
	}

	laddr := &net.TCPAddr{
		Port: (int)(cfg.ServerPort),
	}
	if cfg.ServerIP != "" {
		if laddr.IP = net.ParseIP(cfg.ServerIP); laddr.IP == nil {
			loger.Errorf("Cannot parse server IP, listening on localhost")
		}
	}

	server := NewProxyServer()
	if server.Listener, err = net.ListenTCP("tcp", laddr); err != nil {
		loger.Panicf("Cannot listening at [%v]: %v", laddr, err)
	}
	loger.Infof("Server starts listening at [%v]", server.Listener.Addr())

	// parse and connect to proxy
	if cfg.ProxyURL == "" {
		server.Dialer = proxy.Direct
	}else{
		var u *url.URL
		if u, err = url.Parse(cfg.ProxyURL); err != nil {
			loger.Errorf("Cannot parse proxy URL, using default proxy")
		}
		var dialer proxy.Dialer
		if dialer, err = proxy.FromURL(u, nil); err != nil {
			loger.Errorf("Cannot connect to proxy, using default proxy")
		}
		server.Dialer = dialer.(proxy.ContextDialer)
	}

	exitch := make(chan struct{}, 1)

	go func(){
		defer func(){
			exitch <- struct{}{}
		}()
		if err := server.Serve(); err != nil {
			loger.Errorf("Error when serving: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

WAIT:
	select {
	case s := <-sigs:
		loger.Infof("Got signal %s", s.String())
		if s == syscall.SIGHUP {
			loger.Infof("Reloading config...")
			ncfg := loadConfig()
			loger.Infof("Reloading plugins...")
			manager.UnloadAll()
			if _, err := manager.LoadFromDir(scriptpath); err != nil {
				loger.Errorf("Cannot load scripts: %v", err)
			}
			if cfg.ServerIP != ncfg.ServerIP || cfg.ServerPort != ncfg.ServerPort {
				loger.Info("Server address changed, restarting server...")
				timeoutCtx, cancel := context.WithTimeout(context.Background(), 16 * time.Second)
				loger.Warn("Closing server...")
				server.Shutdown(timeoutCtx)
				cancel()
				loger.Info("Restarting server...")
				cfg = ncfg
				goto RESTART
			}
			if cfg.Debug != ncfg.Debug {
				if ncfg.Debug {
					loger.SetLevel(logger.DebugLevel)
				}else{
					loger.SetLevel(logger.InfoLevel)
				}
			}
			cfg = ncfg
			goto WAIT
		}
		loger.Infof("Unloading plugins...")
		manager.UnloadAll()
		timeoutCtx, _ := context.WithTimeout(context.Background(), 16 * time.Second)
		loger.Warn("Closing server...")
		server.Shutdown(timeoutCtx)
	case <-exitch:
		os.Exit(1)
	}
}

type ProxyServer struct {
	ctx    context.Context
	cancel context.CancelFunc

	Dialer   proxy.ContextDialer
	Listener net.Listener

	conns *Set[*liter.Conn]
}

func NewProxyServer()(s *ProxyServer){
	s = &ProxyServer{
		conns: NewSet[*liter.Conn](),
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	return
}

func (s *ProxyServer)Serve()(err error){
	var c net.Conn
	for {
		if c, err = s.Listener.Accept(); err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return
		}
		go s.handle(liter.WrapConn(c))
	}
}

func (s *ProxyServer)Shutdown(ctx context.Context)(err error){
	if err = s.Listener.Close(); err != nil {
		return
	}
	s.cancel()
	select {
	case <-s.conns.AfterEmpty():
		return
	case <-ctx.Done():
		return ctx.Err()
	}
}
