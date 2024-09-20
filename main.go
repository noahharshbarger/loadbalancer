package main

import (
	"loadbalancer/balancer"
	"loadbalancer/config"
	"log"
	"net/http"
)

func main() {
    // Load the configuration
    config.LoadConfig()

    // Define and initialize the backends
    backends := config.GetBackends()

    // Create the load balancer
    lb := balancer.NewLoadBalancer(backends)

    go lb.HealthCheck()

    // Start the server
    http.HandleFunc("/", lb.HandleRequest)
    log.Println("Load balancer running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}