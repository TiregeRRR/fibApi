package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TiregeRRR/fibApi/config"
	f "github.com/TiregeRRR/fibApi/fibonacci"
	"github.com/gorilla/mux"
)

// StartREST добавляет хэндлеры и запускает сервер
func StartREST() {
	cfg := config.GetConfig()
	port := cfg.GetString("rest_port")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/", writeFibonacci).Methods("GET").Queries("x", "{x}").Queries("y", "{y}")
	log.Printf("Starting REST server on %v port", port)
	log.Fatalln(http.ListenAndServe(":"+port, r))
}

// writeFibonacci обрабатывает запрос и вызывает writeResponseError или writeResponseSuccess
func writeFibonacci(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	sl, code, err := f.GetFibSlice(args["x"], args["y"])
	if err != nil {
		writeResponse(w, []byte(err.Error()), code)
		return
	}
	b, err := json.Marshal(sl)
	if err != nil {
		writeResponse(w, []byte(err.Error()), code)
	}
	writeResponse(w, b, code)
}

// writeResponseError пишет в ResponseWriter
func writeResponse(w http.ResponseWriter, data []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
