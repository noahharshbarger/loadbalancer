package balancer

import (
	"loadbalancer/config"
	"net/http"
	"sync"
)

type LoadBalancer struct {
    backends []config.Backend
    mu       sync.Mutex
}

func NewLoadBalancer(backends []config.Backend) *LoadBalancer {
    return &LoadBalancer{
        backends: backends,
    }
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Implement request handling logic
}