package main

import (
	"flag"
	"hash/fnv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	log.Println("=== MAIN STARTED ===")
	service := flag.String("service", "loadbalancer", "Service to run : loadbalancer, backend1, backend2, or backend3")
	flag.Parse()

	log.Printf("=== SERVICE: %s ===\n", *service)

	switch *service {
	case "loadbalancer":
		log.Println("Starting load balancer...")
		startLoadBalancer()

	case "backend1":
		log.Println("Starting backend1...")
		startBackend1()

	case "backend2":
		log.Println("Starting backend2...")
		startBackend2()

	case "backend3":
		log.Println("Starting backend3...")
		startBackend3()

	default:
		log.Fatal("Valid services: loadbalancer, backend1, backend2, backend3")
	}
}

func startLoadBalancer() {
	http.HandleFunc("/", IPHash)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IPHash(w http.ResponseWriter, r *http.Request) {

	temp := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
	hashing := fnv.New32()
	hashing.Write([]byte(temp))

	backend := [3]string{"backend1:8001", "backend2:8002", "backend3:8003"}
	ip_hash := hashing.Sum32()
	target_url, _ := url.Parse("http://" + backend[ip_hash%3])
	proxy := httputil.NewSingleHostReverseProxy(target_url)
	proxy.ServeHTTP(w, r)
}
