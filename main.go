package main

import (
	"syscall"
)

func main() {

	// Creating a socket file descriptor
	// file handle (windows) or file description (UNIX) - a temporary number assigned to a file by operating system
	// AF_NET - address family, the socket will be using ipv4
	// SOCK_STREAM - connection gets established between two parties until terminated
	fileHandle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err.Error())
	}

	socketAddress := &syscall.SockaddrInet4{
		Port: 3001,
		Addr: [4]byte{127,0,0,1},
	}

	// Binding a socket to a port
	if err := syscall.Bind(fileHandle, socketAddress);err != nil {

	}


}