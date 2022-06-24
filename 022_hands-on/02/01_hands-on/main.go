package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	// Accept
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// write to connection
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}

}
