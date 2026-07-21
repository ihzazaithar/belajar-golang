package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	contentType := varAsRequest.Header.Get("content-type")
	fmt.Fprint(varAsWriter, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	varAsWriter.Header().Add("X-Powered-By", "Ihza Zaidan Afthar")
	fmt.Fprint(varAsWriter, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}
