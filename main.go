package main

import (
	"net/http"
	"strings"
	"fmt"
	"log"
)

func serveGreet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprint(w, "Hello, " + path[7:] + "!")		
}

func serveShout(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprint(w, strings.ToUpper("Hello, " + path[7:] + "!"))		
}

func main() {
	http.HandleFunc("/greet/", serveGreet)
	http.HandleFunc("/shout/", serveShout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
