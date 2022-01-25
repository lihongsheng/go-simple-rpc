package common

import (
	"context"
	"net"
)



var ReadHandler  = make([]func(context *LinkHandler,param interface{}) ,0)
var WriteHandler = make( []func(context *LinkHandler,param interface{}), 0)


type LinkHandler struct {
	ReadIndex int
	WriteIndex int
	Ct context.Context
	Con net.Conn
}

// 读取
func (l *LinkHandler) Rnext() func(context *LinkHandler,param interface{}) {
	l.ReadIndex++
	return ReadHandler[l.ReadIndex]
}

// 写入
func (l *LinkHandler) Write(p interface{})  {
	l.WriteIndex = 0
	WriteHandler[l.WriteIndex](l,p)
}
