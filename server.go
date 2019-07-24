package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/urfave/cli"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func StartListener(c *cli.Context) error {

	listeningPort := c.GlobalString("listening-port")

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	staticHandler := http.FileServer(statikFS)

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(staticHandler)

	http.Handle("/", logRequest(metricsCollector.Handler("", router)))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/ready", readyCheck)
	http.HandleFunc("/log-collector", logUserInput)

	log.Printf("Server starting on port %v... \n", listeningPort)
	log.Println("Web Interface: http://localhost:" + listeningPort + "/")
	log.Println("Prometheus Metrics: http://localhost:" + listeningPort + "/metrics")
	log.Println("Liveness Endpoint: http://localhost:" + listeningPort + "/health")
	log.Println("Readiness Endpoint: http://localhost:" + listeningPort + "/ready")

	err = http.ListenAndServe(":"+listeningPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

// Healthcheck endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Up")
}

// Readiness endpoint
func readyCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ready")
}

// Metrics Middleware.
var metricsCollector = middleware.New(middleware.Config{
	Recorder: metrics.NewRecorder(metrics.Config{}),
})

// Logging Middleware
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested URL: %v\n", r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

// Log user input POST request
func logUserInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var input struct {
			Message string `json:"message"`
		}
		err := decoder.Decode(&input)
		if err != nil {
			log.Println(err)
		}
		log.Println("User sent message:", input.Message)
	}
}
