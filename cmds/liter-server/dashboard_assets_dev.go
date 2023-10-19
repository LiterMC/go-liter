
//go:build dev

package main

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var DashboardAssets http.FileSystem = gin.Dir("dashboard/dist", false)

func init(){
	basedir := "."
	if _, curfile, _, ok := runtime.Caller(0); ok {
		basedir = filepath.Dir(curfile)
	}
	loger.Infof("Starting frontend debug server")
	cmd := exec.Command("npm", "-C", filepath.Join(basedir, "dashboard"), "run", "build-dev")
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		loger.Fatalf("Cannot start the frontend server: %v", err)
	}
	go func(){
		defer cmd.Process.Kill()
		if err := cmd.Wait(); err != nil {
			loger.Errorf("npm frontend server exited: %v", err)
		}
	}()
}
