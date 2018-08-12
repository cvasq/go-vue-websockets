package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Set custom port by running with --port PORT_NUM
	// Default port is 3000
	httpPort := flag.String("port", "3000", "HTTP Listening Address")
	flag.Parse()

	// Initialize file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Start the server and log any errors
	log.Println("http server started on port", ":"+*httpPort)
	err := http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
