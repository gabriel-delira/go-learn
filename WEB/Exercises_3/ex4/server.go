package ex4

import (
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8084")
	if err != nil {
		// handle error
	}
	defer lst.Close()

	for {
		conn, err := lst.Accept()
		if err != nil {
			// handle error
		}
		io.WriteString(conn, "I see you connected")
		defer conn.Close()
	}
}
