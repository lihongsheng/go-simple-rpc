package service

import (
	"log"
	"net"
	"strconv"
)
var readHandler  = make([]func(param interface{}) ,0)
var writeHandler = make( []func(param interface{}), 0)

type LinkHandler struct {
	index int,
	ct con
}

type Server struct {
	l net.Listener
	readHandler  []func(param interface{}) interface{}
	writeHandler []func(param interface{})
}

var DefaultServer = new(Server)

func (s *Server) Listen(port int, scheme string) error {
	var err error
	s.l,err = net.Listen(scheme, ":" + strconv.Itoa(port))
	return err
}

func (s *Server) Run() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func handlerSlert()  {
	
}

func  AddReadHandler(func(param interface{}))  {

}

func (s *Server) Close()  {
	
}