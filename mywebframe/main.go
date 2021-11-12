package main

import (
	"net/http"
	"webframe/farmework"
)

func main() {
	core := farmework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Addr:    "8080",
		Handler: core,
	}
	server.ListenAndServe()
}


