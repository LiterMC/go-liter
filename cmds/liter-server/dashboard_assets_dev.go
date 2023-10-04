
//go:build dev

package main

import (
	"net/http"
	"os"
	"os/exec"
)

var DashboardAssets http.Handler = http.FileServer(http.Dir("dashboard/dist"))

func init(){
	loger.Infof("Starting frontend debug server")
	cmd := exec.Command("npm", "-C", "dashboard", "run", "build-dev")
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
