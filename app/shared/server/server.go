package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores the hostname and port number
type Server struct {
	Hostname  string `json:"Hostname"`  // Server name
	Port  int    `json:"Port"`  // HTTP port
}

// Run starts the HTTP and/or HTTPS listener
func Run(httpHandlers http.Handler, s Server) {
	load(httpHandlers, s)
}

// startHTTP starts the HTTP listener
func load(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))
	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
}

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.Port)
}