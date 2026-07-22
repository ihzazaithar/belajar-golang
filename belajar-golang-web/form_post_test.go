package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Handler
func FormPost(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	// Parsing data dari POST method
	err := varAsRequest.ParseForm()
	if err != nil {
		panic(err)
	}
	// Mengambil data jika Parsing Data berhasil
	firstName := varAsRequest.PostForm.Get("firstName")
	lastName := varAsRequest.PostForm.Get("lastName")

	// Memparsing data sekaligus mengambil datanya dalam satu kali
	// firstName := varAsRequest.PostFormValue.Get("firstName")
	// lastName := varAsRequest.PostFormValue.Get("lastName")

	fmt.Fprintf(varAsWriter, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Ihza&lastName=Afthar")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") // Wajib ditambahkan ketika menggunakan method POST
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body) // mengembalikan result dan error

	fmt.Println(string(body))
}
