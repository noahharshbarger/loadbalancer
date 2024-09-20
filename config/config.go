package config

import "fmt"

type Backend struct {
    URL   string
    Alive bool
}

func GetBackends() []Backend {
    return []Backend{
        {URL: "http://localhost:8081", Alive: true},
        {URL: "http://localhost:8082", Alive: true},
    }
}

func LoadConfig() {
    // Placeholder for loading configuration
    fmt.Println("Configuration loaded")
}