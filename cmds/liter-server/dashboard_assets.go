
//go:build !dev

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate npm -C dashboard ci
//go:generate npm -C dashboard run build
//go:embed dashboard/dist
var _dashboardDist embed.FS
var DashboardAssets http.FileSystem = func()(http.FileSystem){
	assetsFS, err := fs.Sub(_dashboardDist, "dashboard/dist")
	if err != nil {
		loger.Panic(err)
	}
	return http.FS(assetsFS)
}()

