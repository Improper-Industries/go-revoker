package main

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/krakendio/bloomfilter/v2/rpc/client"
)

type RevokeRequest struct {
	Key string `json:"key"`
	Subject string `json:"subject"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	bloomServer, err := client.New("localhost:8020")
	if err != nil {
		panic(err)
	}
	defer bloomServer.Close()

	var req RevokeRequest
	err2 := json.NewDecoder(r.Body).Decode(&req)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Adding %s to %s", req.Key, req.Subject)
	subject := req.Key + "-" + req.Subject
	log.Printf("Adding subject: " + subject)
	err = bloomServer.Add([]byte(subject))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", StatusHandler).Methods("GET")
	r.HandleFunc("/check", CheckHandler).Methods("POST")
	r.HandleFunc("/add", AddHandler).Methods("POST")
	srv := &http.Server{
		Addr: ":8008",
		Handler: r,
	}
	srv.ListenAndServe()
}
