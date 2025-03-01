package loadbalancer

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	//"time"
)

type LoadBalancer struct {
	Servers []*Server
	config  *Config
}

var errorTemplate *template.Template

func init() {
	path := filepath.Join("templates", "503.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic("Failed to load 503 template: " + err.Error())
	}
	errorTemplate = tmpl
}

func NewLoadBalancer(config *Config) *LoadBalancer {
	var servers []*Server
	for _, s := range config.Servers {
		u, _ := url.Parse(s)
		servers = append(servers, &Server{URL: u, Healthy: true})
	}
	return &LoadBalancer{
		Servers: servers,
		config:  config,
	}
}

// nextServerLeastActive Algorithm
func (lb *LoadBalancer) nextServerLeastActive() *Server {
	var leastActiveServer *Server
	leastActiveConnections := -1

	for _, server := range lb.Servers {
		server.Mutex.Lock()
		if (server.ActiveConnections < leastActiveConnections || leastActiveConnections == -1) && server.Healthy {
			leastActiveConnections = server.ActiveConnections
			leastActiveServer = server
		}
		server.Mutex.Unlock()
	}

	return leastActiveServer
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path, "Server:", r.Host)
	server := lb.nextServerLeastActive()
	if server == nil {
		log.Println("No healthy servers available")
		w.WriteHeader(http.StatusServiceUnavailable)
		errorTemplate.Execute(w, nil)
		return
	}

	server.Mutex.Lock()
	server.ActiveConnections++
	server.Mutex.Unlock()
	server.Proxy().ServeHTTP(w, r)
	server.Mutex.Lock()
	//time.Sleep(3 * time.Second)
	server.ActiveConnections--
	server.Mutex.Unlock()

	log.Println("Request completed:", r.Method, r.URL.Path, "Server:", r.Host, "Server URL:", server.URL)
}
