
//go:build !dev

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate npm -C dashboard run build
//go:embed dashboard/dist
var _dashboardDist embed.FS
var DashboardAssets http.Handler = func()(http.Handler){
	assetsFS, err := fs.Sub(_dashboardDist, "dashboard/dist")
	if err != nil {
		loger.Panic(err)
	}
	return http.FileServer(http.FS(assetsFS))
}()

