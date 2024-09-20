package balancer

import (
	"loadbalancer/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestHealthCheck ensures that the health check mechanism is working as expected.
func TestHealthCheck(t *testing.T) {
	// Create a healthy and an unhealthy backend server
	healthyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer healthyServer.Close()

	unhealthyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer unhealthyServer.Close()

	// Simulate backend servers
	backends := []config.Backend{
		{URL: healthyServer.URL, Alive: true},
		{URL: unhealthyServer.URL, Alive: true},
	}

	lb := NewLoadBalancer(backends)

	// Start the health check
	go lb.HealthCheck()

	// Wait for health check to update statuses
	time.Sleep(2 * time.Second)

	// Check that the first server is still healthy
	if !lb.backends[0].Alive {
		t.Error("Expected first server to be alive")
	}

	// Check that the second server is now unhealthy
	if lb.backends[1].Alive {
		t.Error("Expected second server to be unhealthy")
	}
}