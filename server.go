package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/middleware"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"

	"github.com/urfave/cli"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func StartListener(c *cli.Context) error {

	listeningPort := c.GlobalString("listening-port")

	// Create our middleware.
	metricsCollector := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(statikFS)

	r := mux.NewRouter()

	r.Handle("/", metricsCollector.Handler("", staticHandler))

	// Serves up the index.html file regardless of the path.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		staticHandler.ServeHTTP(w, r)
	})

	http.Handle("/static/", staticHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Server starting on port %v... \n", listeningPort)
	log.Println("Web Interface: http://localhost:" + listeningPort + "/")
	log.Println("Prometheus Metrics: http://localhost:" + listeningPort + "/metrics")

	err = http.ListenAndServe(":"+listeningPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}
