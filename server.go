package main

import (
	"fmt"
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

	// Create our middleware.
	metricsCollector := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(statikFS)

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(staticHandler)

	http.Handle("/", logRequest(metricsCollector.Handler("", router)))
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/log-collector", logUserInput)

	log.Printf("Server starting on port %v... \n", listeningPort)
	log.Println("Web Interface: http://localhost:" + listeningPort + "/")
	log.Println("Prometheus Metrics: http://localhost:" + listeningPort + "/metrics")

	err = http.ListenAndServe(":"+listeningPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested URL: %v\n", r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})

}

func logUserInput(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		fmt.Println(r.URL.Query())

		keys, ok := r.URL.Query()["data"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'data' is missing")
			return
		}

		// Query()["key"] will return an array of items,
		// we only want the single item.
		data := keys[0]

		log.Println("Url Param 'data' is: " + string(data))

	}
}
