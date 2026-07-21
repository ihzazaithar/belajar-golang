package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) { // writer adalah Pesan atau sesuatu yang akan ditampilkan ke client, sedangkan request adalah apapun yang dikirim oleh client ke server dan ditangkap oleh handler.
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	varAsMux := http.NewServeMux()

	varAsMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	varAsMux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, World")
	})
	varAsMux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images") // semua url yang setelah / terakhir pada /images/ maka akan merujuk ke sini. berbeda dengan /images, hanya unik ke images ketika memasukan ke /images
	})
	varAsMux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnails") // Lebih spesifik/panjang menang. ketika kita masuk ke url ini maka akan masuk ke /images/thumbnail bukan ke /images/
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: varAsMux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
