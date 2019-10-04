package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	defer c.Close()
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Received: " + netData)
	c.Write([]byte("You said: " + netData))
}

func main() {
	l, err := net.Listen("tcp4", ":5002")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
