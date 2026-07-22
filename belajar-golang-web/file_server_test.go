package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// Implementasi Static FileServer
func TestFileServer(t *testing.T) {
	varAsDir := http.Dir("./resources") // Persiapkan Folder yang ingin dijadikan Static File Server

	varAsFileServer := http.FileServer(varAsDir)

	mux := http.NewServeMux()
	// mux.Handle("/static/", varAsFileServer) // Ini menghasilkan 404 not found, karena FileServer akan mencari file dengan path /resources/static/namaFile. sedangkan, file path berada pada /resources/namaFile
	mux.Handle("/static/", http.StripPrefix("/static/", varAsFileServer)) // Untuk mengatasi 404 not found, gunakan http.StripPrefix("prefix_url_yang_ingin_dihilangkan", varAsFileServer)

	varAsServer := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := varAsServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Integrasi Golang Embed dengan FileServer
//
//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	varAsDir, _ := fs.Sub(resources, "resources") // Ini untuk mengatasi masalah 404 not found. Dengan ini program masuk ke "subdir" dari static folder file server
	varAsFileServer := http.FileServer(http.FS(varAsDir))

	varAsMux := http.NewServeMux()
	varAsMux.Handle("/static/", http.StripPrefix("/static", varAsFileServer))

	varAsServer := http.Server{
		Addr:    "localhost:8080",
		Handler: varAsMux,
	}

	err := varAsServer.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
