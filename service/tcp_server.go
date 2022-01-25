package service

import (
	"context"
	"go-simple-rpc/common"
	"log"
	"net"
	"strconv"

)

type Server struct {
	l net.Listener
}

var ServerList = make(map[string]interface{})

func RegisterServer(serverName string,server interface{})  {
	ServerList[serverName] = server
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
		go s.handlerConn(conn)
	}
}

func (s *Server) handlerConn(conn net.Conn)  {
	for  {
		// 建立请求处理链和上下文
		// 这个上下未加入超时控制
		ct := &common.LinkHandler{
			ReadIndex: 0,
			WriteIndex: 0,
			Ct: context.Background(),
			Con: conn,
		}
		common.ReadHandler[ct.ReadIndex](ct,conn)
	}
}

func  AddReadHandler(f func(context *common.LinkHandler,param interface{}))  {
	common.ReadHandler = append(common.ReadHandler,f)
}

func  AddWriteHandler(f func(context *common.LinkHandler,param interface{}))  {
	common.WriteHandler = append(common.WriteHandler,f)
}

func (s *Server) Close()  {
	
}