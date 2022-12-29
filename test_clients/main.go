package main

import (
	"fmt"
	"net"
)

func main() {
	c,err := net.Dial("tcp", "127.0.0.1:3001")
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			panic("error closing connection: " + err.Error())
		}
	}(c)
	if err != nil {
		panic(err.Error())
	}
	lenn, err := c.Write([]byte("Hello there"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("writing len: ", lenn)
}
