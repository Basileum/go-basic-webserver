package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHelloHandlerGetOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8082/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestHelloHandlerOtherNOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8082/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if want, got := http.StatusNotFound, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestHelloHandlerWrongRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8082/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if want, got := http.StatusNotFound, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestFormHandler(t *testing.T) {
	form := url.Values{}
	form.Add("name", "name_example")
	form.Add("address", "adress_example")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8082/form", strings.NewReader(form.Encode()))
	w := httptest.NewRecorder()

	formHandler(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
