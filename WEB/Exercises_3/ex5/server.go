package ex5

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8085")
	if err != nil {
		// handle error
	}
	defer lst.Close()

	for {
		conn, err := lst.Accept()
		if err != nil {
			// handle error
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Printf("Received %s\n", ln)
		}
		defer conn.Close()

		fmt.Println("Code got here. Nerver gets Here because the scanner dont know when to stop reading the stream")
		io.WriteString(conn, "I see you connected")
	}
}
