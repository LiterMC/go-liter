
package main

import (
	"net/http"
)

var _ http.Handler = (*Server)(nil)

func (s *Server)initHandler(){
	mux := http.NewServeMux()
	mux.Handle("/main/", http.StripPrefix("/main", DashboardAssets))
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request){
		if req.URL.Path != "/" {
			http.NotFound(rw, req)
			return
		}
		http.Redirect(rw, req, "/main", http.StatusFound)
	})
	s.handler = mux
}

func (s *Server)ServeHTTP(rw http.ResponseWriter, req *http.Request){
	s.handler.ServeHTTP(rw, req)
}
