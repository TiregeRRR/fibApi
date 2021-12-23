package rest

import (
	"encoding/json"
	"log"
	"net/http"

	f "interview/fibonacci"

	"github.com/gorilla/mux"
)

func StartREST() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	log.Fatalln(http.ListenAndServe(":8000", r))
}

func writeFibonacci(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	sl, err := f.GetFibSlice(args["x"], args["y"])
	if err != nil {
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(sl)
	if err != nil {
		writeResponseError(w, err, http.StatusInternalServerError)
	}
	writeResponseSuccess(w, b, http.StatusOK)
}

func writeResponseError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(err.Error()))
}

func writeResponseSuccess(w http.ResponseWriter, data []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
