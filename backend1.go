package main

import (
	"log"
	"net/http"
)

func startBackend1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend server 1"))
	})

	log.Println("Starting Backend 1 on 0.0.0.0:8001")
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", mux))
}
