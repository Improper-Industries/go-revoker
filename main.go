package main

import (
	"os"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": 200, "message": "OK"}`))
}

type check struct {
	Status int `json:"status"`
	Subject string `json:"subject"`
	Exists bool `json:"exists"`
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	bloomServer, err := client.New(os.Getenv("BLOOM_SERVER"))
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

	subject := req.Key + "-" + req.Subject

	exists, _ := bloomServer.Check([]byte(subject))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(check{
		Status: 200,
		Subject: subject,
		Exists: exists,
	})
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	bloomServer, err := client.New(os.Getenv("BLOOM_SERVER"))
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
	subject := req.Key + "-" + req.Subject
	err = bloomServer.Add([]byte(subject))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Added:", subject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status": 201, "message": "Added"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", StatusHandler).Methods("GET")
	r.HandleFunc("/check", CheckHandler).Methods("POST")
	r.HandleFunc("/add", AddHandler).Methods("POST")
	srv := &http.Server{
		Addr: ":3005",
		Handler: r,
	}
	log.Println("Starting Revoker on port 3005")
	log.Fatal(srv.ListenAndServe())
}
