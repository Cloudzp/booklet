package server

import (
	"github.com/jeffallen/mqtt"
	"net"
)

type Server struct {
	addr string
}

func NewServer(addr string) (*Server, error){
	s:= &Server{
		addr:addr,
	}
  return s,nil
}

func (s *Server)Run() error {
     li,err := net.Listen("tcp",s.addr)
     if err != nil {
     	return err
	 }
     server := mqtt.NewServer(li)
     server.Start()

     <- server.Done
     return nil
}