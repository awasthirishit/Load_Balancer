package main

import (
	"fmt"
	"net/http"
	"sync"
)

// holds info for available servers and the current index.
type LoadBalancer struct {
	servers []string
	current int
	mutex   sync.Mutex
}

// this function creates a new LoadBalancer instance with provided server addresses.
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		current: 0,
	}
}

// selecting server using round-robin manner.
func (lb *LoadBalancer) ChooseServer() string {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

func main() {

	serverList := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	// Creating a load balancer instance
	lb := NewLoadBalancer(serverList)

	// handler function for load balancing requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		server := lb.ChooseServer()
		fmt.Fprintf(w, "Load Balancer routed request to server: %s\n", server)
	}

	// Set up a HTTP server
	http.HandleFunc("/", handler)
	fmt.Println("Load Balancer listening on port 9090...")
	http.ListenAndServe(":9090", nil)
}
