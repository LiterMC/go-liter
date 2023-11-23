package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/kmcsr/go-liter/script"
)

func parseFlags() {
	var command string
	if len(os.Args) == 1 {
		command = "serve"
	} else {
		command = os.Args[1]
		switch command {
		case "daemon":
			if len(os.Args) < 3 {
				loger.Fatal("Missing argument <pidfile>, need point a pid file for the daemon process")
			}
			proca := &os.ProcAttr{
				Files: []*os.File{nil, os.Stdout, os.Stderr},
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
				Files: []*os.File{nil, os.Stdout, os.Stderr},
			}
			args := make([]string, len(os.Args)-1)
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
				loger.Fatal("Missing argument <pidfile>, need point a pid file for the daemon process")
			}
			pidfile = os.Args[2]
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
}

func main() {
	parseFlags()

	cleaner, err := runDashResources()
	if err != nil {
		loger.Panicf("Cannot start the frontend server: %v", err)
	}
	defer cleaner()

	r := new(Runner)
	if r.configDir, err = os.Getwd(); err != nil {
		loger.Panicf("Cannot get working dir", err)
	}
	r.scriptPath = filepath.Join(r.configDir, "plugins")
	r.cfgHash, _ = genRandB64(48)
	r.config = loadConfig()
	r.whitelist = loadWhitelist()
	r.blacklist = loadBlacklist()
	r.manager = script.NewManager()
	r.server = NewServer(r.configDir, r.manager)
	r.server.configLock = &r.configLock
	r.server.cfgHash = &r.cfgHash
	r.server.config = &r.config
	r.server.blacklist = &r.blacklist
	r.server.whitelist = &r.whitelist

	r.startServer()
	if r.config.Dashboard.Enable {
		r.startDashboard(r.config.Dashboard.Addr)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

WAIT:
	select {
	case s := <-sigs:
		loger.Warnf("Got signal %s", s.String())
		if s == syscall.SIGHUP { // reload
			r.Reload()
			goto WAIT
		}
		r.Stop(sigs)
	case <-r.exitSvr:
		os.Exit(1)
	case <-r.exitDash:
		// ignore the http server exit
		r.exitDash = nil
		goto WAIT
	}
	loger.Warn("Server exited")
}

type Runner struct {
	configDir  string
	scriptPath string

	configLock sync.RWMutex
	cfgHash    string
	config     Config
	whitelist  Whitelist
	blacklist  Blacklist

	manager           *script.Manager
	server            *Server
	dashboard         *http.Server
	exitSvr, exitDash chan struct{}
}

func (r *Runner) Run() {
	var err error

	loger.Infof("Liter Server %s", version)
	r.manager.SetLogger(loger)

	if _, err := os.Stat(r.scriptPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(r.scriptPath, 0755)
	}

	if _, err = r.manager.LoadFromDir(r.scriptPath); err != nil {
		loger.Errorf("Cannot load scripts: %v", err)
	}

	return
}

func (r *Runner) startServer() {
	if r.exitSvr != nil {
		panic("liter: Server is already started")
	}
	r.server.Addr = r.config.ServerAddr
	listener, err := net.Listen("tcp", r.server.Addr)
	if err != nil {
		loger.Panicf("Cannot listening at [%s] for liter server: %v", r.server.Addr, err)
		return
	}
	loger.Infof("Liter server listening at [%v]", listener.Addr())
	exit := make(chan struct{}, 0)
	go func() {
		defer close(exit)
		if err := r.server.Serve(listener); err != nil {
			loger.Errorf("Error on serve: %v", err)
		}
	}()
	r.exitSvr = exit
}

func (r *Runner) startDashboard(addr string) {
	if r.exitDash != nil || r.dashboard != nil {
		panic("liter: Dashboard is already started")
	}
	r.dashboard = &http.Server{
		Addr:    addr,
		Handler: r.server,
	}
	if r.server.users.Len() == 0 {
		loger.Errorf("No any users were found, creating one...")
		if passwd, err := genRandB64(16); err != nil {
			loger.Errorf("Cannot create new user: %v", err)
		} else {
			root := &User{
				Name: "root",
			}
			// one more sha from the browser
			root.SetPassword(asSha256Hex(passwd))
			if err := r.server.users.AddUser(root); err != nil {
				loger.Errorf("Cannot create new user: %v", err)
			} else {
				loger.Infof("Root user created: password=%s", passwd)
			}
		}
	}
	var err error
	if r.exitDash, err = listenAndServeHTTP(r.dashboard); err != nil {
		loger.Panicf("Cannot listening at [%s] for http server: %v", r.dashboard.Addr, err)
	}
}

func (r *Runner) closeServer(ctx context.Context) {
	loger.Warn("Closing server ...")
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	}
	r.server.Shutdown(ctx)
	r.exitSvr = nil
}

func (r *Runner) closeDashboard(ctx context.Context) {
	loger.Warn("Closing dashboard ...")
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	}
	r.dashboard.Shutdown(ctx)
	r.dashboard = nil
	r.exitDash = nil
}

func (r *Runner) ReloadConfigs() (old Config) {
	r.configLock.Lock()
	defer r.configLock.Unlock()

	r.cfgHash, _ = genRandB64(48)
	old = r.config
	r.config = loadConfig()
	r.whitelist = loadWhitelist()
	r.blacklist = loadBlacklist()
	return
}

func (r *Runner) Reload() {
	loger.Info("Unloading plugins ...")
	r.manager.UnloadAll()

	loger.Infof("Reloading config...")
	ocfg := r.ReloadConfigs()

	if ocfg.ServerAddr != r.config.ServerAddr {
		loger.Info("Restarting ...")
		r.startServer()
	}
	if !r.config.Dashboard.Enable {
		// if disabled but dashboard is opening
		if r.dashboard != nil {
			r.closeDashboard(nil)
		}
	} else if r.dashboard == nil {
		// if enabled but dashboard is closed
		loger.Info("Starting dashboard ...")
		r.startDashboard(r.config.Dashboard.Addr)
	} else if ocfg.Dashboard.Enable && ocfg.Dashboard.Addr != r.config.Dashboard.Addr {
		// if enabled but address changed
		r.closeDashboard(nil)
		loger.Info("Restarting dashboard ...")
		r.startDashboard(r.config.Dashboard.Addr)
	}

	loger.Info("Reloading plugins ...")
	if _, err := r.manager.LoadFromDir(r.scriptPath); err != nil {
		loger.Errorf("Cannot load plugins: %v", err)
	}
}

func (r *Runner) Stop(sigs <-chan os.Signal) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func() {
		defer cancel()
		select {
		case s := <-sigs:
			loger.Errorf("Got second close signal %s", s.String())
		}
	}()

	done := make(chan struct{}, 2)
	hasDash := r.dashboard != nil

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		r.closeServer(timeoutCtx)
	}()
	if hasDash {
		go func() {
			defer func() {
				done <- struct{}{}
			}()
			r.closeDashboard(timeoutCtx)
		}()
	}
	<-done
	if hasDash {
		<-done
	}
}

func listenAndServeHTTP(server *http.Server) (exited chan struct{}, err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.Addr); err != nil {
		return
	}
	loger.Infof("HTTP server listening at [%v]", listener.Addr())
	exit := make(chan struct{}, 0)
	go func() {
		defer close(exit)
		if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			loger.Errorf("Error on serve: %v", err)
		}
	}()
	return exit, nil
}
