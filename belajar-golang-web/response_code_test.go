package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	name := varAsRequest.URL.Query().Get("name")

	if name == "" {
		// varAsWriter.WriteHeader(http.StatusBadRequest) // dengan mengirim variable
		varAsWriter.WriteHeader(400) // dengan mengirim nilai
		fmt.Fprint(varAsWriter, "Name is Empty")
	} else {
		// varAsWriter.WriteHeader(200) // Tidak perlu dideklarasikan karena defaultnya akan mengirim ini.
		fmt.Fprintf(varAsWriter, "Hello %s", name)
	}
}

func TestReponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
}

func TestReponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ihza", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
}
