package ex8

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8088")
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
		go serve(conn)

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

}

func write(conn net.Conn, value string) {
	io.WriteString(conn, value)
}
