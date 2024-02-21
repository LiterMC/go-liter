//go:build dev

package main

import (
	"bufio"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var DashboardAssets http.FileSystem

func runDashResources() (cleaner func(), err error) {
	basedir := "."
	if _, curfile, _, ok := runtime.Caller(0); ok {
		basedir = filepath.Dir(curfile)
	}

	loger.Infof("Starting frontend debug server")
	npmDir := filepath.Join(basedir, "dashboard")
	distDir := filepath.Join(npmDir, "dist")

	DashboardAssets = gin.Dir(distDir, false)

	cmd := exec.Command("npm", "-C", npmDir, "run", "build-dev")

	pout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	perr, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	go func() {
		s := bufio.NewScanner(pout)
		for s.Scan() {
			loger.Info("[npm-dev/stdout]: " + s.Text())
		}
	}()
	go func() {
		s := bufio.NewScanner(perr)
		for s.Scan() {
			loger.Info("[npm-dev/stderr]: " + s.Text())
		}
	}()

	go func() {
		if err := cmd.Wait(); err != nil {
			loger.Errorf("npm frontend server exited: %v", err)
		}
	}()

	cleaner = func() {
		cmd.Process.Kill()
	}
	return
}
