package main

import (
	"fmt"
	"github.com/nenadmitt/tcp"
	"io"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	s,err := tcp.NewSocket(3004)
	defer s.Close()
	if err != nil {
		panic(err)
	}

	for {
		c,err := s.AcceptConnection()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("READING CONECCTION")
		bytes,err := io.ReadAll(c)
		fmt.Println(string(bytes))
	}
}