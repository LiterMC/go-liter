
package main

import (
	"context"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter"
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
	var (
		err error
		server = NewServer(nil)
		cfg = getConfig()
	)

	RESTART:
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
			handler(liter.NewConn(c))
			s.cond.L.Lock()
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
			s.cond.L.Unlock()
		}(c)
	}
}

func (s *Server)Shutdown(ctx context.Context)(err error){
	if err = s.Listener.Close(); err != nil {
		return
	}
	select {
		case <-waitUntilNot(s.cond, func()(bool){ return len(s.conns) > 0 }):
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
	loger.Debugf("client [%v] connected", c.RemoteAddr())
	var (
		p *liter.PacketReader
		err error
	)
	if p, err = c.Recv(); err != nil {
		loger.Errorf("client [%v] read handshake packet error: %v", c.RemoteAddr(), err)
		return
	}
	var hp liter.HandshakeP
	if hp, err = liter.ReadHandshakeP(p); err != nil {
		loger.Errorf("client [%v] read handshake packet error: %v", c.RemoteAddr(), err)
		return
	}
	loger.Tracef("client [%v] handshake packet: %v", c.RemoteAddr(), hp)

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
		loger.Infof("client [%v] try connected unexcept address %q", c.RemoteAddr(), hp.Addr)
		return
	}

	loger.Infof("client [%v] connected with address [%s:%d], passing to server '%s'", c.RemoteAddr(), hp.Addr, hp.Port, svr.Id)

	if hp.NextState == liter.NextPingState && svr.HandlePing {
		loger.Debugf("handle ping connection [%v] for server '%s'", c.RemoteAddr(), svr.Id)
		if p, err = c.Recv(); err != nil {
			loger.Errorf("client [%v] read ping request packet error: %v", c.RemoteAddr(), err)
			return
		}
		if err = liter.ReadPingRequestP(p); err != nil {
			loger.Errorf("client [%v] read ping request packet error: %v", c.RemoteAddr(), err)
		}
		c.Send(0x00, liter.Object{
			"version": liter.Object{
				"name": "Idle",
				"protocol": 0,
			},
			"players": liter.Object{
				"max": 1,
				"online": 0,
			},
			"description": liter.Object{
				"text": svr.Motd,
			},
		})
		return
	}

	var conn net.Conn
	if conn, err = net.Dial("tcp", svr.Target); err != nil {
		loger.Errorf("Cannot dial to %q: %v", svr.Target, err)
		return
	}
	np := liter.NewPacket(p.Id())
	hp.Encode(np)
	if _, err = conn.Write(np.Bytes()); err != nil {
		loger.Errorf("New connection handshake error: %v", err)
		return
	}
	rc := c.RawConn()
	go io.Copy(rc, conn)
	io.Copy(conn, rc)
}
