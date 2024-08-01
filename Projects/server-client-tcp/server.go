package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1245")
	if err != nil {
		fmt.Errorf("ERROR:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Listening on port 1245")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("ERROR:", err.Error())
		}
		go HandleConnection(conn)

	}

}

func HandleConnection(conn net.Conn) {
	buff := make([]byte, 4096)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println(string(buff))
	conn.Close()

}
