package ex11

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

// Server - Creates my tcp server
func Server() {
	lst, err := net.Listen("tcp", ":8091")
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

	var method, url string
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			splitedLine := strings.Split(line, " ")
			method = splitedLine[0]
			url = splitedLine[1]
		}
		if line == "" {
			break
		}
		fmt.Println("Received: \t", line)
		counter++
	}

	values := []string{
		"CHECK OUT THE RESPONSE BODY PAYLOAD:",
		fmt.Sprintf("Method => %s", method),
		fmt.Sprintf("URL => %s", url),
	}
	write(conn, values...)
}

func write(conn net.Conn, values ...string) {
	body := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
		<h1>HOLY COW THIS IS LOW LEVEL</h1>
		<h2>%s</h2>
		<p>%s</p>
		<p>%s</p>
	</html>`, values[0], values[1], values[2])

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", body)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	io.WriteString(conn, body)

	fmt.Println("Response written")
}
