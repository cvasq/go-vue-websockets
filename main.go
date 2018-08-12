package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Set custom port by running with --port PORT_NUM
	// Default port is 8080
	httpPort := flag.String("port", "8080", "HTTP Listening Address")
	flag.Parse()

	// Initialize file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Start the server and log any errors
	log.Println("HTTP server started on port:", *httpPort)
	err := http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
