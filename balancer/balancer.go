package balancer

import (
	"loadbalancer/config"
	"net/http"
	"sync"
)

type LoadBalancer struct {
    backends []config.Backend
    mu       sync.Mutex
    current  int
}

func NewLoadBalancer(backends []config.Backend) *LoadBalancer {
    return &LoadBalancer{
        backends: backends,
        current:  0,
    }
}

func (lb *LoadBalancer) GetNextBackend() config.Backend {
    lb.mu.Lock()
    defer lb.mu.Unlock()

    backend := lb.backends[lb.current]
    lb.current = (lb.current + 1) % len(lb.backends)
    return backend
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Implement request handling logic
}