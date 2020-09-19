package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// request(conn)

	url := getURL(conn)
	fmt.Println("Recovered", url)
	if url == "/" {
		responseDefault(conn)
	}
	if url == "/test" {
		responseTest(conn)
	}
}

func request(conn net.Conn) {
	counter := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if counter == 0 {
			requestStringsList := strings.Fields(currentLine)
			method := requestStringsList[0]
			url := requestStringsList[1]
			fmt.Printf(" ***Method: %s\n ***URL: %s\n", method, url)
		}
		// End of received message
		if currentLine == "" {
			break
		}
		counter++
	}
}

func getURL(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	url := ""
	for scanner.Scan() {
		line := scanner.Text()
		url = strings.Fields(line)[1]
		break
	}
	return url
}

func responseDefault(conn net.Conn) {
	buf := new(bytes.Buffer)
	tpl.Execute(buf, "")

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", buf.Len())
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	err := tpl.Execute(conn, "")
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

func responseTest(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
		<h1>Here is my mutiplex test working guys !!!</h1>
	</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", body)
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
