package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port     = flag.String("port", "", "Port to start the web server on.")
	filename = flag.String("file", "index.html", "File that contains the form data you'd like to test.")
)

func init() {
	flag.StringVar(port, "p", "", "Port to start the web server on")
	flag.StringVar(filename, "f", "index.html", "File that contains the form data you'd like to test.")
}

func usage() {
	fmt.Println("ERROR: Missing Argument")
	fmt.Println("Usage:")
	fmt.Println("    ./autofill-pwn -p 8080 [-f index.html]")
	os.Exit(1)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fileToServe, err := ioutil.ReadFile(*filename)
		if err != nil {
			log.Printf("Error opening file: %s", err)
		} else {
			fmt.Fprintf(w, string(fileToServe))
			log.Printf("Serving GET Request...")
		}
	case "POST":
		log.Println("Serving POST request! Data incoming!")
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form %s", err)
		} else {
			for k, v := range r.Form {
				log.Printf("Field %s = %s", k, v)
			}
		}
	}
}

func main() {
	flag.Parse()
	log.Printf("Starting server on port %s...", *port)
	if len(*port) == 0 {
		usage()
	}
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":"+*port, nil)
}
