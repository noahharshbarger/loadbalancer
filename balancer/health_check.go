package balancer

import (
	"net/http"
	"time"
)

func (lb *LoadBalancer) HealthCheck() {
    for {
        for i, backend := range lb.backends {
            resp, err := http.Get(backend.URL)
            if err != nil || resp.StatusCode != http.StatusOK {
                lb.backends[i].Alive = false
            } else {
                lb.backends[i].Alive = true
            }
        }
        time.Sleep(30 * time.Second) // Adjust the interval as needed
    }
}