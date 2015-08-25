package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

var port = flag.Int("port", 8900, "port number")

func main() {
	xxx()
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func xxx() {
	var f, err = os.OpenFile("dummy.txt", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0660)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("bla")
	defer f.Close()
}
