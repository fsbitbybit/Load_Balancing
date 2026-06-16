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

	log.Fatal(http.ListenAndServe(":8002", mux))
}
