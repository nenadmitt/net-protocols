package tcp

import (
	"fmt"
	"os"
	"syscall"
)

type socket struct {
	fileHandle syscall.Handle
}

func NewSocket(port int) (*socket, error) {

	// Creating a socket file descriptor
	// file handle (windows) or file description (UNIX) - a temporary number assigned to a file by operating system
	// AF_NET - address family, the socket will be using ipv4
	// SOCK_STREAM - connection gets established between two parties until terminated
	fileHandle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, os.NewSyscallError("socket", err)
	}

	socketAddress := &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{127,0,0,1},
	}

	if err := syscall.Bind(fileHandle, socketAddress); err != nil {
		return nil, os.NewSyscallError("bind", err)
	}

	if err := syscall.Listen(fileHandle, syscall.SOMAXCONN); err != nil {
		return nil, os.NewSyscallError("listen", err)
	}

	return &socket{fileHandle: fileHandle}, nil
}

func (s *socket) AcceptConnection() (*socket, error){
	nfd, addr, err := syscall.Accept(s.fileHandle)
	if err == nil {
		syscall.CloseOnExec(nfd)
	}
	fmt.Println("Acception connection")
	fmt.Println(nfd, " FILE HANDLE")
	fmt.Println(addr, " ADDR")
	fmt.Println(err, " ERR")
	if err != nil {
		return nil, err
	}
	return &socket{nfd}, nil
	//if err == nil {
	//	syscall.CloseOnExec(nfd)
	//}
}

func (s *socket) Read(data []byte) (int, error) {
	n, err := syscall.Read(s.fileHandle, data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *socket) Close() error {
	return syscall.Close(s.fileHandle)
}
