package main

import (
	"hash/fnv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func main() {
	go startBackend1()
	go startBackend2()
	go startBackend3()

	time.Sleep(1 * time.Second)

	http.HandleFunc("/", IPHash)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IPHash(w http.ResponseWriter, r *http.Request) {

	temp := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
	hashing := fnv.New32()
	hashing.Write([]byte(temp))

	backend := [3]string{"8001", "8002", "8003"}
	ip_hash := hashing.Sum32()
	target_url, _ := url.Parse("http://localhost:" + backend[ip_hash%3])
	proxy := httputil.NewSingleHostReverseProxy(target_url)
	proxy.ServeHTTP(w, r)
}
