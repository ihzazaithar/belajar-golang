package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=ihza", nil)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Ihza&last_name=Afthar", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(varAsWriter http.ResponseWriter, varAsRequest *http.Request) {
	var varAsQuery url.Values = varAsRequest.URL.Query()
	var varAsMap []string = varAsQuery["names"]

	fmt.Fprint(varAsWriter, strings.Join(varAsMap, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?names=Ihza&names=Zaidan&names=Afthar", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
