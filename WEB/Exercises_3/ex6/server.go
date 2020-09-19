package ex6

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Server() {
	lst, err := net.Listen("tcp", ":8086")
	if err != nil {
		fmt.Println("ERROR opening connection", err.Error())
	}
	defer lst.Close()

	for {
		conn, err := lst.Accept()
		if err != nil {
			fmt.Println("ERROR accepting request", err.Error())
		}

		// READ
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				break
			}
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}

}
