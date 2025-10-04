package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("listening on port :6379")

	//starting the server
	l, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println(err)
		return
	}

	// listening to server
	con, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer con.Close()

	for {
		buf := make([]byte, 1024)

		//reading message from client
		_, err := con.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client", err.Error())
			os.Exit(1)
		}

		con.Write([]byte("+OK\r\n"))
	}
}
