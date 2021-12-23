package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var writeFibonacciExpected = "[0,1,1,2,3,5]"

func TestWriteFibonacciBadRequest1(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/?x=-1&y=5", nil)
	w := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	m.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v got %v", http.StatusBadRequest, res.StatusCode)
	}
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "invalid input: x < 0" {
		t.Errorf("expected response %v got %v", "invalid input: x < 0", string(data))
	}
}

func TestWriteFibonacciBadRequest2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/?x=1&y=-5", nil)
	w := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	m.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v got %v", http.StatusBadRequest, res.StatusCode)
	}
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "invalid input: y < 0" {
		t.Errorf("expected response %v got %v", "invalid input: y < 0", string(data))
	}
}
func TestWriteFibonacciBadRequest3(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/?x=55&y=5", nil)
	w := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	m.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v got %v", http.StatusBadRequest, res.StatusCode)
	}
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != "invalid input: x > y" {
		t.Errorf("expected response %v got %v", "invalid input: x > y", string(data))
	}
}

func TestWriteFibonacciOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8000/api/v1/?x=0&y=5", nil)
	w := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	m.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v got %v", http.StatusOK, res.StatusCode)
	}
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != writeFibonacciExpected {
		t.Errorf("expected response %v got %v", writeFibonacciExpected, string(data))
	}
}
