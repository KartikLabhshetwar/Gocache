package main

import (
	"fmt"
	"net"
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
		resp := NewResp(con)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		//ignore request and send back a ok request
		con.Write([]byte("+OK\r\n"))
	}
}
