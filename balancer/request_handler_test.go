// request_handler_test.go
package balancer

import (
	"loadbalancer/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRequestHandling ensures that requests are forwarded to the correct backend server
func TestRequestHandling(t *testing.T) {
    // Create a backend server
    backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer backendServer.Close()

    // Simulate a backend
    backends := []config.Backend{
        {URL: backendServer.URL, Alive: true},
    }

    lb := NewLoadBalancer(backends)

    // Create a request to the load balancer
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    // Handle the request with the load balancer
    lb.HandleRequest(w, req)

    // Check that the response is OK
    if status := w.Result().StatusCode; status != http.StatusOK {
        t.Errorf("Expected status 200, got %d", status)
    }
}
