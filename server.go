package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli"
)

func StartListener(c *cli.Context) error {

	listeningPort := c.GlobalString("listening-port")

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(FileServer)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server listening on port", listeningPort)
	log.Println("Access the web UI at http://localhost:" + listeningPort + "/")
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
