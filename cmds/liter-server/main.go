
package main

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
	// "github.com/kmcsr/go-liter/packets"
)

var _after_load = initBeforeLoad()

func initBeforeLoad()(_ bool){
	var command string
	if len(os.Args) == 1 {
		command = "serve"
	}else{
		command = os.Args[1]
		switch command {
		case "daemon":
			if len(os.Args) < 3 {
				os.Args = append(os.Args, "/var/run/litermc.pid")
			}
			proca := &os.ProcAttr{
				Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			}
			os.Args[1] = "__daemon0"
			name, err := os.Executable()
			if err != nil {
				loger.Fatalf("Cannot get the executable: %v", err)
			}
			proc, err := os.StartProcess(name, os.Args, proca)
			if err != nil {
				loger.Fatalf("Cannot start daemon0 process: %v", err)
			}
			loger.Infof("Daemon0 process %d started", proc.Pid)
			stat, err := proc.Wait()
			if err != nil {
				loger.Fatalf("Daemon0 process error: %v", err)
			}
			ecode := stat.ExitCode()
			loger.Infof("Daemon0 exited with status %d", ecode)
			os.Exit(ecode)
		case "__daemon0":
			if len(os.Args) < 3 {
				loger.Fatal("Missing argument <pidfile>, need point a pid file for the daemon process")
			}
			pidfile := os.Args[2]
			proca := &os.ProcAttr{
				Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			}
			args := make([]string, len(os.Args) - 1)
			args[0] = os.Args[0]
			args[1] = "serve"
			copy(args[2:], os.Args[3:])
			name, err := os.Executable()
			if err != nil {
				loger.Fatalf("Cannot get the executable: %v", err)
			}
			proc, err := os.StartProcess(name, args, proca)
			if err != nil {
				loger.Fatalf("Cannot start daemon process: %v", err)
			}
			loger.Infof("Daemon process %d started", proc.Pid)
			if err := os.WriteFile(pidfile, ([]byte)(strconv.Itoa(proc.Pid)), 0644); err != nil {
				loger.Errorf("Cannot write pid file '%s': %v", pidfile, err)
			}
			os.Exit(0)
		case "reload":
			var pidfile string
			if len(os.Args) < 3 {
				pidfile = "/var/run/litermc.pid"
			}else{
				pidfile = os.Args[2]
			}
			buf, err := os.ReadFile(pidfile)
			if err != nil {
				loger.Fatalf("Cannot read pid file '%s': %v", pidfile, err)
			}
			pid, err := strconv.Atoi((string)(buf))
			if err != nil {
				loger.Fatalf("Cannot parse pid file content %q: %v", (string)(buf), err)
			}
			if pid < 0 {
				loger.Fatalf("Pid must larger than 0, got %d", pid)
			}
			proc, err := os.FindProcess(pid)
			if err != nil {
				loger.Fatalf("Cannot find process by pid %d: %v", pid, err)
			}
			if err = proc.Signal(syscall.SIGHUP); err != nil {
				loger.Fatalf("Cannot send SIGHUP to process %d: %v", pid, err)
			}
			os.Exit(0)
		}
	}
	if command == "serve" {
		return
	}
	loger.Fatalf("Unknown subcommand '%s'", command)
	return
}

func main(){
	scriptpath := filepath.Join(configDir, "plugins")
	var (
		err error
		server = NewServer(nil)
		cfg = getConfig()
		manager = script.NewManager()
	)

	if cfg.Debug {
		loger.SetLevel(logger.TraceLevel)
	}else{
		loger.SetLevel(logger.InfoLevel)
	}

	loger.Infof("Liter Server %s", version)
	manager.SetLogger(loger)

RESTART:
	if _, err := os.Stat(scriptpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(scriptpath, 0755)
	}
	if _, err = manager.LoadFromDir(scriptpath); err != nil {
		loger.Errorf("Cannot load scripts: %v", err)
	}

	laddr := &net.TCPAddr{
		Port: (int)(cfg.ServerPort),
	}
	if len(cfg.ServerIP) > 0 {
		if laddr.IP = net.ParseIP(cfg.ServerIP); laddr.IP == nil {
			loger.Warnf("Cannot parse server-ip")
		}
	}
	if server.Listener, err = net.ListenTCP("tcp", laddr); err != nil {
		loger.Fatalf("Cannot listening at [%v]: %v", laddr, err)
	}
	loger.Infof("Server listening at [%v]", server.Listener.Addr())

	exitch := make(chan struct{}, 1)

	go func(){
		defer func(){
			exitch <- struct{}{}
		}()
		if err := server.Serve(); err != nil {
			loger.Errorf("Error on serve: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

WAIT:
	select {
	case s := <-sigs:
		loger.Infof("Got signal %s", s.String())
		if s == syscall.SIGHUP {
			reloadConfigs()
			ncfg := getConfig()
			loger.Info("Reloading plugins ...")
			manager.UnloadAll()
			if _, err = manager.LoadFromDir(scriptpath); err != nil {
				loger.Errorf("Cannot load scripts: %v", err)
			}
			if cfg.Debug != ncfg.Debug {
				if ncfg.Debug {
					loger.SetLevel(logger.TraceLevel)
				}else{
					loger.SetLevel(logger.InfoLevel)
				}
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
			cfg = ncfg
			goto WAIT
		}
		timeoutCtx, _ := context.WithTimeout(context.Background(), 16 * time.Second)
		loger.Warn("Closing server...")
		server.Shutdown(timeoutCtx)
	case <-exitch:
		os.Exit(1)
	}
}

type Server struct{
	Listener net.Listener
	conns    []net.Conn
	cond     *sync.Cond
}

func NewServer(listener net.Listener)(s *Server){
	return &Server{
		Listener: listener,
		cond: sync.NewCond(new(sync.Mutex)),
	}
}

func (s *Server)Serve()(err error){
	var c net.Conn
	for {
		if c, err = s.Listener.Accept(); err != nil {
			return
		}
		go func(c net.Conn){
			s.cond.L.Lock()
			s.conns = append(s.conns, c)
			s.cond.L.Unlock()
			defer func(){
				s.cond.L.Lock()
				defer s.cond.L.Unlock()
				e := len(s.conns) - 1
				for i, d := range s.conns {
					if d == c {
						if i != e {
							s.conns[i], s.conns[e] = s.conns[e], s.conns[i]
						}
						s.conns = s.conns[:e]
						if len(s.conns) == 0 {
							s.cond.Broadcast()
						}
						break
					}
				}
			}()
			handler(liter.WrapConn(c))
		}(c)
	}
}

func (s *Server)Shutdown(ctx context.Context)(err error){
	if err = s.Listener.Close(); err != nil {
		return
	}
	select {
	case <-waitUntilNot(s.cond, func()(bool){
		return len(s.conns) > 0
	}):
	case <-ctx.Done():
		err = ctx.Err()
		s.cond.L.Lock()
		if len(s.conns) > 0 {
			for _, c := range s.conns {
				c.Close()
			}
			s.conns = nil
			s.cond.Broadcast()
		}
		s.cond.L.Unlock()
	}
	return
}

func handler(c *liter.Conn){
	defer c.Close()
	ploger := logger.NewPrefixLogger(loger, "client [%v]:", c.RemoteAddr())
	ploger.Debugf("Connected!")
	var err error
	var hp *liter.HandshakePkt
	if hp, err = c.RecvHandshakePkt(); err != nil {
		ploger.Errorf("read handshake packet error: %v", err)
		return
	}
	ploger.Tracef("Handshake packet: %v", hp)

	var svr *ServerIns = nil
	servers := getServers()
	F: for _, s := range servers {
		for _, n := range s.ServerNames {
			if ismatch(hp.Addr, n) {
				svr = s
				break F
			}
		}
	}
	if svr == nil {
		ploger.Infof("Trying connected unexcept address %q", hp.Addr)
		return
	}

	ploger.Infof("Connected with address [%s:%d], passing to server %q[%s]", hp.Addr, hp.Port, svr.Id, svr.Target)

	if hp.NextState == liter.NextPingState && svr.HandlePing {
		ploger.Debugf("Handle ping connection for server %q", svr.Id)
		handleServerStatus(ploger, c, "Idle", svr.Motd)
		return
	}

	var conn *liter.Conn
	if conn, err = liter.Dial(svr.Target); err != nil {
		ploger.Errorf("Cannot dial to %q: %v", svr.Target, err)
		ploger.Debugf("Handle ping connection for server %q", svr.Id)
		handleServerStatus(ploger, c, "Closed", svr.MotdFailed)
		return
	}
	ploger.Debugf("Target %q connected", svr.Id)
	if err = conn.SendHandshakePkt(hp); err != nil {
		ploger.Errorf("New connection handshake error: %v", err)
		return
	}
	rc := c.RawConn()
	go io.Copy(rc, conn.RawConn())
	io.Copy(conn.RawConn(), rc)
}

func handleServerStatus(loger logger.Logger, c *liter.Conn, name string, motd string){
	var srp liter.StatusRequestPkt
	var err error
	if err = c.RecvPkt(0x00, &srp); err != nil {
		loger.Errorf("Read status request packet error: %v", err)
		return
	}
	if err = c.SendPkt(0x00, liter.StatusResponsePkt{
		Payload: liter.Object{
			"version": liter.Object{
				"name": name,
				"protocol": c.Protocol(),
			},
			"players": liter.Object{
				"max": 1,
				"online": 0,
			},
			"description": liter.Object{
				"text": motd,
			},
		},
	}); err != nil {
		loger.Errorf("Send packet error: %v", err)
		return
	}
	var prp liter.PingRequestPkt
	if err = c.RecvPkt(0x01, &prp); err != nil {
		loger.Errorf("Read ping request packet error: %v", err)
		return
	}
	if err = c.SendPkt(0x01, (liter.PingResponsePkt)(prp)); err != nil {
		loger.Errorf("Send ping response packet error: %v", err)
		return
	}
}
