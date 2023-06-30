package main

import (
	"github.com/upupnoah/hade/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: framework.NewCore(), // custom request core handler
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
