package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Server struct {
	URL               *url.URL
	ActiveConnections int
	Healthy           bool
	Mutex             sync.Mutex
}

func (s *Server) Proxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(s.URL)
}

func (s *Server) CheckHealth() {
	res, err := http.Get(s.URL.String() + "/health")
	if err != nil {
		log.Printf("Health check failed for %s: %v", s.URL.Host, err)
		s.Healthy = false
		return
	}
	defer res.Body.Close()

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if res.StatusCode != http.StatusOK {
		log.Printf("Server %s is unhealthy (Status: %d)", s.URL.Host, res.StatusCode)
		s.Healthy = false
	} else {
		log.Printf("Server %s is healthy", s.URL.Host)
		s.Healthy = true
	}
}

func StartHealthChecks(servers []*Server, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, server := range servers {
			go server.CheckHealth()
		}
	}
}
