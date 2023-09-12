package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{"/", http.StatusNotFound, "404 Not Found"},
		{"/?non-valid-url", http.StatusBadRequest, "400 Bad Request"},
		{"/?https://www.google.com\" onload=\"alert('XSS')", http.StatusBadRequest, "400 Bad Request"},
		{"/?https://www.google.com", http.StatusOK, "Redirecting to https://www.google.com"},
		{"/?https%3A%2F%2Fwww.google.com%2Fsearch%3Fq%3Dgolang", http.StatusOK, "Redirecting to https://www.google.com/search?q=golang"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectedStatus)
		}

		if strings.Contains(rr.Body.String(), test.expectedBody) == false {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.expectedBody)
		}
	}
}
