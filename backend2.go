package main

import (
	"log"
	"net/http"
)

func startBackend2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend server 2"))
	})

	log.Println("Starting Backend 2 on 0.0.0.0:8002")
	log.Fatal(http.ListenAndServe("0.0.0.0:8002", mux))
}
