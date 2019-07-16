package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/urfave/cli"
)

func StartListener(c *cli.Context) error {

	listeningPort := c.GlobalString("listening-port")

	// Create our middleware.
	metricsCollector := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(FileServer)

	http.Handle("/", metricsCollector.Handler("", router))
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Server starting on port %v... \n", listeningPort)
	log.Println("Web Interface: http://localhost:" + listeningPort + "/")
	log.Println("Prometheus Metrics: http://localhost:" + listeningPort + "/metrics")
	err := http.ListenAndServe(":"+listeningPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

// Serve web files in public directory
func FileServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	extension, _ := regexp.MatchString("\\.+[a-zA-Z]+", r.URL.EscapedPath())
	// If the url contains an extension, use file server
	if extension {
		http.FileServer(http.Dir("./websocket-echo-client/dist")).ServeHTTP(w, r)
	} else {
		http.ServeFile(w, r, "./websocket-echo-client/dist/index.html")
	}
}
