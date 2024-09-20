package balancer

import (
	"loadbalancer/config"
	"testing"
)

// TestRoundRobin ensures requests are being distributed to the backend servers in a round-robin fashion.
func TestRoundRobin(t *testing.T) {
    // Simulate backend servers
    backends := []config.Backend{
        {URL: "http://localhost:8081", Alive: true},
        {URL: "http://localhost:8082", Alive: true},
    }

    lb := NewLoadBalancer(backends)

    // Check that the first request goes to the first server
    if backend := lb.GetNextBackend(); backend.URL != "http://localhost:8081" {
        t.Errorf("Expected http://localhost:8081, got %s", backend.URL)
    }

    // Check that the second request goes to the second server
    if backend := lb.GetNextBackend(); backend.URL != "http://localhost:8082" {
        t.Errorf("Expected http://localhost:8082, got %s", backend.URL)
    }

    // Check that the third request goes back to the first server
    if backend := lb.GetNextBackend(); backend.URL != "http://localhost:8081" {
        t.Errorf("Expected http://localhost:8081, got %s", backend.URL)
    }
}