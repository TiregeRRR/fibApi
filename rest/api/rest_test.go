package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var testCases = []struct {
	x, y           int
	expectedStatus int
	expectedBody   string
}{
	{
		x:              -1,
		y:              5,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   "invalid input: x < 0",
	},
	{
		x:              1,
		y:              -5,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   "invalid input: y < 0",
	},
	{
		x:              -1,
		y:              5,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   "invalid input: x < 0",
	},
	{
		x:              5,
		y:              1,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   "invalid input: x > y",
	},
	{
		x:              5,
		y:              1,
		expectedStatus: http.StatusOK,
		expectedBody:   "[0,1,1,2,3,5]",
	},
}

func TestRESTWriteFibonacciBadRequest(t *testing.T) {
	for _, v := range testCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/?x=%v&y=%v", v.x, v.y), nil)
		w := httptest.NewRecorder()
		m := mux.NewRouter()
		m.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
		m.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != v.expectedStatus {
			t.Errorf("expected status %v got %v", http.StatusBadRequest, res.StatusCode)
		}
		data, _ := ioutil.ReadAll(res.Body)
		if string(data) != v.expectedBody {
			t.Errorf("expected response %v got %v", "invalid input: x < 0", string(data))
		}
	}
}
