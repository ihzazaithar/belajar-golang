package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Handler membuat cookie
func SetCookie(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	varAsCookie := new(http.Cookie)
	varAsCookie.Name = "X-PZN-Name" // Nama untuk Cookie yang dibuat
	varAsCookie.Value = varAsRequest.URL.Query().Get("name")
	varAsCookie.Path = "/"

	http.SetCookie(varAsWriter, varAsCookie)
	fmt.Fprint(varAsWriter, "Success Create Cookie")
}

// Handler mengambil cookie
func GetCookie(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	varAsCookie, err := varAsRequest.Cookie("X-PZN-Name") // Mengembalikan Result dan Error
	if err != nil {
		fmt.Fprint(varAsWriter, "No Cookie")
	} else {
		name := varAsCookie.Value
		fmt.Fprintf(varAsWriter, "Hello %s", name)
	}
}

// Live server
func TestCookie(t *testing.T) {
	varAsMux := http.NewServeMux()
	varAsMux.HandleFunc("/set-cookie", SetCookie)
	varAsMux.HandleFunc("/get-cookie", GetCookie)

	varAsServer := http.Server{
		Addr:    "localhost:8080",
		Handler: varAsMux,
	}

	err := varAsServer.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// Unit Test Web untuk membuat Cookie
func TestSetCookie(t *testing.T) {
	varAsRequest := httptest.NewRequest(http.MethodGet, "localhost:8080?name=Ihza", nil)
	varAsRecorder := httptest.NewRecorder()

	SetCookie(varAsRecorder, varAsRequest)

	cookies := varAsRecorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Value, cookie.Name)
	}
}

// Unit Test Web untuk mengambil Cookie
func TestGetCookie(t *testing.T) {
	varAsRequest := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	varAsCookie := new(http.Cookie)
	varAsCookie.Name = "X-PZN-Name"
	varAsCookie.Value = "Ihza"
	varAsRequest.AddCookie(varAsCookie)

	varAsRecorder := httptest.NewRecorder()

	GetCookie(varAsRecorder, varAsRequest)

	response, _ := io.ReadAll(varAsRecorder.Result().Body)
	fmt.Println(string(response))
}
