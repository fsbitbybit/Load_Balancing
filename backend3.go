package main

import (
	"log"
	"net/http"
)

func startBackend3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend server 3"))
	})

	log.Println("Starting Backend 3 on 0.0.0.0:8003")
	log.Fatal(http.ListenAndServe("0.0.0.0:8003", mux))
}
