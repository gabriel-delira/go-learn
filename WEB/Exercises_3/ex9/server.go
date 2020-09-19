package ex9

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8089")
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
		fmt.Println("Connected")

		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println("Received: \t", line)
	}

	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	write(conn, body)
}

func write(conn net.Conn, body string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	fmt.Println("Response written")
}
