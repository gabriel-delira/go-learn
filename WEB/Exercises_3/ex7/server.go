package ex7

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8087")
	if err != nil {
		fmt.Println("Error creating server")
	}
	defer lst.Close()

	for {
		conn, err := lst.Accept()
		if err != nil {
			fmt.Println("Error connecting to server")
			continue
		}

		go serve(conn)

		fmt.Println("Connected")
		io.WriteString(conn, "Connection accepted")
		conn.Close()
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println("Received: \t", line)
	}
}
