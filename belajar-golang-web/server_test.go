package belajar_golang_web

import (
	"net/http"
	"testing"
)

// ** Membuat Web Server

func TestServer(t *testing.T) {
	varAsServer := http.Server{
		Addr: "localhost:8080",
	}

	varErr := varAsServer.ListenAndServe()
	if varErr != nil {
		panic(varErr)
	}

	// Jalankan Programnya go test -v -run=TestServer atau go run namaFile
	// dan buka localhost:8080 di web browser.
}

// Membuat Web Server *
