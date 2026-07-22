package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	if varAsRequest.URL.Query().Get("name") != "" {
		http.ServeFile(varAsWriter, varAsRequest, "./resources/ok.html")
	} else {
		http.ServeFile(varAsWriter, varAsRequest, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	varAsServer := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := varAsServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Implementasi Embed sebagai ServeFile
//
//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	if varAsRequest.URL.Query().Get("name") != "" {
		fmt.Fprint(varAsWriter, resourceOk)
	} else {
		fmt.Fprint(varAsWriter, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	varAsServer := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := varAsServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
