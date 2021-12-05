package service

import (
	"net"
	"strconv"
)

type Server struct {
	l net.Listener
	readHandler  []func(param interface{})
	writeHandler []func(param interface{})
}

var DefaultServer = new(Server)

func (s *Server) Listen(port int, scheme string) error {
	var err error
	s.l,err = net.Listen(scheme, ":" + strconv.Itoa(port))
	return err
}

func (s *Server) AddReadHandler(func(param interface{}))  {

}

func (s *Server) Close()  {
	
}