package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

func main() {

	// Set custom port by running with --port PORT_NUM
	// Default port is 8080
	httpPort := flag.String("port", "8080", "HTTP Listening Address")
	flag.Parse()

	log.Println("Starting Server")

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(FileServer)
	http.Handle("/", r)

	log.Println("Listening on port: ", *httpPort)
	err := http.ListenAndServe(":"+*httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Serve web files in public directory
func FileServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	extension, _ := regexp.MatchString("\\.+[a-zA-Z]+", r.URL.EscapedPath())
	// If the url contains an extension, use file server
	if extension {
		http.FileServer(http.Dir("./public/")).ServeHTTP(w, r)
	} else {
		http.ServeFile(w, r, "./public/index.html")
	}
}
