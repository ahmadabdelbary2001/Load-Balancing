package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	statusEndpoint = "/status"
	homeEndpoint   = "/"
	htmlFile       = "templates/index.html"
)

type WebServer struct {
	Port       int
	ServerName string
	Template   *template.Template
	Mu         sync.Mutex
}

func NewWebServer(port int, name string) *WebServer {
	ws := &WebServer{
		Port:       port,
		ServerName: name,
	}
	ws.LoadTemplate()
	return ws
}

func (ws *WebServer) LoadTemplate() {
	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}
	ws.Template = tmpl
}

func (ws *WebServer) Start() {
	http.HandleFunc(homeEndpoint, ws.HandleHome)
	http.HandleFunc(statusEndpoint, ws.HandleStatus)

	log.Printf("Server %s starting on :%d", ws.ServerName, ws.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ws.Port), nil))
}

func (ws *WebServer) HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	data := struct{ ServerName string }{ws.ServerName}

	ws.Mu.Lock()
	defer ws.Mu.Unlock()
	if err := ws.Template.Execute(w, data); err != nil {
		log.Printf("Template error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ws *WebServer) HandleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Server is alive\n")
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <port> <server-name>")
	}

	port := os.Args[1]
	name := os.Args[2]

	var portNum int
	if _, err := fmt.Sscanf(port, "%d", &portNum); err != nil {
		log.Fatal("Invalid port number")
	}

	server := NewWebServer(portNum, name)
	server.Start()
}
