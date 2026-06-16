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

	log.Fatal(http.ListenAndServe(":8003", mux))
}
