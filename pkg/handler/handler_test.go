package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	long_url  = "https://codesubmit.io/library/react"
	short_rul = "http://short.est/XVlBzgbaiC"
)

// should return error response when wrong method provided
func TestEnocdeURLMethodError(t *testing.T) {
	input := URL{URL: long_url}
	expected := `{"message":"Method not allowed"}`
	req, err := http.NewRequest("GET", input.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EnocdeURL)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// should return success response url encoded
func TestEnocdeURL(t *testing.T) {
	input := URL{URL: long_url}
	expected := `{"short_url":"http://short.est/XVlBzgbaiC","message":"URL encoded successfully"}`
	req, err := http.NewRequest("POST", "/encode", strings.NewReader(input.URL))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EnocdeURL)
	handler.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// should return not found error response whe wrong short url provided or not available
func TestDecodeURLError(t *testing.T) {
	input := URL{URL: "some-wrong-url"}
	expected := `{"message":"URL not found"}`
	req, err := http.NewRequest("POST", "/decode", strings.NewReader(input.URL))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DecodeURL)
	handler.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// should return error response when wrong method provided
func TestDecodeURLMethodError(t *testing.T) {
	input := URL{URL: long_url}
	expected := `{"message":"Method not allowed"}`
	req, err := http.NewRequest("PUT", input.URL, strings.NewReader(input.URL))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DecodeURL)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
