
# Go HTTP Load Balancer

A simple HTTP load balancer built with Go that demonstrates techniques like round-robin load balancing, health checking for backend services, and concurrent request handling. 

## Features
- **Round-robin load balancing**: Distributes requests evenly across multiple backend servers.
- **Health checking**: Periodically checks the availability of backend servers and only forwards requests to healthy servers.
- **Concurrency**: Utilizes Go's goroutines to handle multiple requests concurrently and to run health checks in the background.

## Prerequisites
- **Go**: Version 1.16 or higher
- **Backends**: The load balancer requires backend services running on different ports (e.g., `localhost:8081`, `localhost:8082`).

## Project Structure
```
loadbalancer/
├── main.go               # Entry point of the application
├── balancer/
│   ├── balancer.go       # Load balancing logic
│   └── health_check.go   # Health checking logic for backends
├── config/
│   └── config.go         # Backend configuration (list of backend services)
├── go.mod                # Go module file
└── README.md             # Project documentation
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/loadbalancer.git
   cd loadbalancer
   ```

2. Initialize Go modules:
   ```bash
   go mod tidy
   ```

3. Ensure that your backends are running on different ports. For testing, you can use Python’s simple HTTP server:
   ```bash
   python3 -m http.server 8081
   python3 -m http.server 8082
   ```

## Usage

1. **Run the load balancer**:
   ```bash
   go run main.go
   ```

   The load balancer will start listening for requests on port `8080`.

2. **Test the load balancer**:
   You can use `curl` to send requests and observe that they are forwarded to the backend servers in a round-robin fashion:

   ```bash
   curl http://localhost:8080
   ```

   Each request should be routed to one of the backend servers (`localhost:8081` or `localhost:8082`), and the output will change depending on which server is handling the request.

## Configuration

The backend servers are configured in the `config/config.go` file. You can add or modify backend URLs in the `GetBackends` function:

```go
func GetBackends() []Backend {
    return []Backend{
        {URL: "http://localhost:8081", Alive: true},
        {URL: "http://localhost:8082", Alive: true},
    }
}
```

## Health Check

The load balancer automatically checks the health of the backend servers every 10 seconds. If a backend server becomes unreachable, it will be marked as unavailable, and no traffic will be forwarded to it until it becomes healthy again.

## Extending the Load Balancer

1. **Additional Load Balancing Strategies**:
   You can implement other algorithms like **least connections** or **IP-hash** by modifying the logic in `balancer/balancer.go`.

2. **TLS/HTTPS Support**:
   To add support for HTTPS, you can modify the `http.ListenAndServe` function in `main.go` to use `http.ListenAndServeTLS` with a valid certificate and private key.

3. **Logging and Monitoring**:
   Add logging to track requests and monitor the health of backend services.

4. **Retry Mechanism**:
   Implement a retry logic for failed requests to improve the robustness of your load balancer.

## License
This project is open-source and available under the [MIT License](LICENSE).
