# Go Load Balancer

A simple yet efficient HTTP load balancer written in Go, designed to distribute incoming traffic across multiple backend servers using the **Least Connections** algorithm. Includes health checks and a custom 503 error page for handling server outages.

---

## Features
- **Least Connections Algorithm**: Distributes requests to the server with the fewest active connections.
- **Health Checks**: Periodically verifies backend server health.
- **Custom 503 Page**: Returns a user-friendly error page when all servers are unavailable.
- **Concurrent Handling**: Efficiently manages multiple simultaneous requests.
- **Configurable**: Easy setup via `config.json`.

---

## Prerequisites
- Go 1.23.2 or later

---

## Project Structure
```
go-load-balancing/
├── config.json # Load balancer configuration
├── main.go # Entry point for the load balancer
├── go.mod # Go module file
├── loadbalancer/
│ ├── config.go # Configuration parsing
│ ├── healthcheck.go # Health check logic
│ └── loadbalancer.go # Core load balancing logic
└── templates/
  └── 503.html # Custom 503 error page
```

---

## Configuration
Modify config.json to customize the load balancer:
```
{
    "listenPort": ":9090",
    "healthCheckInterval": "5s",
    "servers": [
        "http://localhost:9091",
        "http://localhost:9092",
        "http://localhost:9093"
    ]
}
```
listenPort: Port for the load balancer (e.g., :9090).
healthCheckInterval: Frequency of health checks (e.g., 5s).
servers: List of backend server URLs.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/ahmadabdelbary2001/Load-Balancing.git
   cd Load-Balancing/go-load-balancing
   ```

2. Build and run:
   ```bash
   go run main.go
   ```

---

## Running the Backend Servers
For the load balancer to work properly, you need to run at least one backend server. We provide two sample implementations:

### Option 1: Go Web Application
1. Navigate to project directory
2. Run with command:
```bash
go run main.go [PORT] "[SERVER_NAME]"
```

### Option 2: Java Web Application
1. Navigate to project directory
2. Ensure you have Java JRE installed
3. Run with command:
```bash
java -jar .\out\artifacts\simpleWebApp_jar\simpleWebApp.jar [PORT] "[SERVER_NAME]"
```

## Important Notes:
1. The load balancer will automatically detect healthy servers
2. If no servers are running, you'll see the 503 error page
3. You can mix Go and Java servers in your configuration
