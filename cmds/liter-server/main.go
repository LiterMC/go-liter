
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
	"syscall"
	"time"

	"github.com/kmcsr/go-liter/script"
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
	return
}

func main(){
	scriptpath := filepath.Join(configDir, "plugins")
	var (
		err error
		cfg, _ = getConfig()
		manager = script.NewManager()
		server = NewServer(manager)
		dashboard *http.Server

		exit1, exit2 chan struct{}
	)

	loger.Infof("Liter Server %s", version)
	manager.SetLogger(loger)

	if _, err := os.Stat(scriptpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(scriptpath, 0755)
	}
	if _, err = manager.LoadFromDir(scriptpath); err != nil {
		loger.Errorf("Cannot load scripts: %v", err)
	}

	startServer := func(addr string){
		server.Addr = addr
		if exit1, err = listenAndServeLiterServer(server); err != nil {
			loger.Fatalf("Cannot listening at [%s] for liter server: %v", cfg.ServerAddr, err)
		}
	}
	startDashboard := func(addr string){
		dashboard = &http.Server{
			Addr: addr,
			Handler: server,
		}
		if server.users.Len() == 0 {
			loger.Errorf("No any users were found, creating one...")
			if passwd, err := genRandB64(16); err != nil {
				loger.Errorf("Cannot create new user: %v", err)
			}else{
				root := &User{
					Name: "root",
				}
				// one more sha from the browser
				root.SetPassword(asSha256Hex(passwd))
				if err := server.users.AddUser(root); err != nil {
					loger.Errorf("Cannot create new user: %v", err)
				}else{
					loger.Infof("Root user created: password=%s", passwd)
				}
			}
		}
		if exit2, err = listenAndServeHTTP(dashboard); err != nil {
			loger.Fatalf("Cannot listening at [%s] for http server: %v", dashboard.Addr, err)
		}
	}

	startServer(cfg.ServerAddr)

	if cfg.Dashboard.Enable {
		startDashboard(cfg.Dashboard.Addr)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

WAIT:
	select {
	case s := <-sigs:
		loger.Warnf("Got signal %s", s.String())
		if s == syscall.SIGHUP {
			reloadConfigs()
			ncfg, _ := getConfig()
			loger.Info("Reloading plugins ...")
			manager.UnloadAll()
			if _, err = manager.LoadFromDir(scriptpath); err != nil {
				loger.Errorf("Cannot load scripts: %v", err)
			}
			if cfg.ServerAddr != ncfg.ServerAddr {
				timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
				loger.Warn("Closing server for restart ...")
				server.Shutdown(timeoutCtx)
				cancel()
				loger.Info("Restarting ...")
				startServer(ncfg.ServerAddr)
			}
			if !ncfg.Dashboard.Enable {
				// if disabled but dashboard is opening
				if dashboard != nil {
					loger.Warn("Closing dashboard ...")
					timeoutCtx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
					dashboard.Shutdown(timeoutCtx)
					dashboard = nil
					cancel()
				}
			}else if dashboard == nil {
				// if enabled but dashboard is closed
				loger.Info("Starting dashboard ...")
				startDashboard(ncfg.Dashboard.Addr)
			}else if cfg.Dashboard.Enable && cfg.Dashboard.Addr != ncfg.Dashboard.Addr {
				// if enabled but address changed
				loger.Warn("Closing dashboard for restart ...")
				timeoutCtx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
				dashboard.Shutdown(timeoutCtx)
				cancel()
				loger.Info("Restarting dashboard ...")
				startDashboard(ncfg.Dashboard.Addr)
			}
			cfg = ncfg
			goto WAIT
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		go func(){
			defer cancel()
			select {
			case <-sigs:
				loger.Errorf("Got second close signal %s", s.String())
			}
		}()
		loger.Warn("Closing server...")
		done := make(chan struct{}, 2)
		go func(){
			defer func(){
				done <- struct{}{}
			}()
			server.Shutdown(timeoutCtx)
		}()
		if dashboard != nil {
			go func(){
				defer func(){
					done <- struct{}{}
				}()
				dashboard.Shutdown(timeoutCtx)
			}()
		}
		<-done
		if dashboard != nil {
			<-done
		}
	case <-exit1:
		os.Exit(1)
	case <-exit2:
		// ignore the http server exit
		exit2 = nil
		goto WAIT
	}
	loger.Warn("Server exited")
}


func listenAndServeLiterServer(server *Server)(exited chan struct{}, err error){
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.Addr); err != nil {
		return
	}
	loger.Infof("Liter server listening at [%v]", listener.Addr())
	exit := make(chan struct{}, 0)
	go func(){
		defer close(exit)
		if err := server.Serve(listener); err != nil {
			loger.Errorf("Error on serve: %v", err)
		}
	}()
	return exit, nil
}

func listenAndServeHTTP(server *http.Server)(exited chan struct{}, err error){
	var listener net.Listener
	if listener, err = net.Listen("tcp", server.Addr); err != nil {
		return
	}
	loger.Infof("HTTP server listening at [%v]", listener.Addr())
	exit := make(chan struct{}, 0)
	go func(){
		defer close(exit)
		if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			loger.Errorf("Error on serve: %v", err)
		}
	}()
	return exit, nil
}
