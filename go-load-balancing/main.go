// في ملف main.go
package main

import (
	"go-load-balancing/loadbalancer"
	"log"
	"net/http"
	"time"
)

func main() {
	config, err := loadbalancer.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	healthCheckInterval, err := time.ParseDuration(config.HealthCheckInterval)
	if err != nil {
		log.Fatalf("Invalid health check interval: %v", err)
	}

	lb := loadbalancer.NewLoadBalancer(config)

	go loadbalancer.StartHealthChecks(lb.Servers, healthCheckInterval)

	http.HandleFunc("/", lb.HandleRequest)

	log.Println("Starting server on port", config.ListenPort)
	err = http.ListenAndServe(config.ListenPort, nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
