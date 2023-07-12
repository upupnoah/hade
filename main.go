package main

import (
	"github.com/upupnoah/hade/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
