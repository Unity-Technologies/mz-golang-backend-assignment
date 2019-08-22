package server

import (
	"fmt"
	"net"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (server *Server) Start(laddr *net.TCPAddr) error {
	fmt.Println("TODO: Start handling client connections and messages")
	return nil
}

func (server *Server) ListClientIDs() []uint64 {
	fmt.Println("TODO: Return the IDs of the connected clients")
	return []uint64{}
}

func (server *Server) Stop() error {
	fmt.Println("TODO: Stop accepting connections and close the existing ones")
	return nil
}
