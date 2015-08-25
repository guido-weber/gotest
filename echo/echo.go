package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var addr = flag.String("addr", ":9100", "Local address to listen on")

func main() {
	flag.Parse()
	l, err := net.Listen("tcp4", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
