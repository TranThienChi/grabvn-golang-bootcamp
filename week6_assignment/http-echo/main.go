package main

import (
	"math/rand"
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Create a new statsd client
	statsdClient, err := statsd.New("127.0.0.1:8080",
		statsd.WithNamespace("counter."),               // prefix every metric with the app name
		statsd.WithTags([]string{"region:us-east-1a"}), // send the EC2 availability zone as a tag with every metric
		// add more options here...
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create our server
	logger := log.New()
	server := Server{
		logger:       logger,
		statsdClient: statsdClient,
	}

	// Start the server
	server.ListenAndServe()
}

// Server represents our server.
type Server struct {
	logger       *log.Logger
	statsdClient *statsd.Client
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() {
	s.logger.Info("echo server is starting on port 8080...")
	http.HandleFunc("/", s.echo)
	http.ListenAndServe(":8080", nil)
}

// Echo echos back the request as a response
func (s *Server) echo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// 30% chance of failure
	if rand.Intn(100) < 30 {
		// add logging on failure
		s.logger.Error("a chaos monkey broke your server")
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		s.statsdClient.Count("fail", 1, nil, 1)
		return
	}

	// Happy path
	writer.WriteHeader(200)
	request.Write(writer)
	s.statsdClient.Count("success", 1, nil, 1)
}
