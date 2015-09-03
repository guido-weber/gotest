package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/satori/go.uuid"
)

var port = flag.Int("port", 8900, "port number")

func main() {
	fmt.Println(uuid.NewV4().String())
	readConfig().printStore()
	/*
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
	*/
}

// Vault is it
type Vault struct {
	Path string
}

// Store has it
type Store struct {
	Name   string
	Age    int
	Vaults []*Vault
}

func (s *Store) printStore() {
	fmt.Printf("%#v\n", s)
	for _, p := range s.Vaults {
		fmt.Printf("  %#v\n", p)
	}
}

func readConfig() *Store {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var s Store
	dec := json.NewDecoder(f)
	dec.Decode(&s)
	return &s
}

func xxx() {
	var f, err = os.OpenFile("dummy.txt", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("bla")
}
