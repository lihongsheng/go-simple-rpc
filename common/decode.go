package common

import (
	"encoding/binary"
	"go-simple-rpc/common/protocol"
	"net"
)

func Decode(ct *LinkHandler,param interface{}) {
	reader := param.(net.Conn)
	h := make([]byte,18)
	//var err error
	_,err := reader.Read(h)
	if err != nil {
		return
	}

	p := protocol.Protocol{}
	// binary.BigEndian.Uint32(h[0:4])
	p.SetVersion(int(binary.BigEndian.Uint32(h[0:4])))
	pl := int(binary.BigEndian.Uint32(h[4:8]))
	p.SetRequestId(int64(binary.BigEndian.Uint64(h[8:16])))
	p.SetRequestType(int(h[16]))
	p.SetBodyType(int(h[17]))
	messageLength := pl - 8 - 1 - 1

	message := make([]byte,messageLength)
	_,err = reader.Read(message)
	if err != nil {
		return
	}
	p.SetBody(message)
	// 下一个链接的处理
	ct.Rnext()(ct,&p)
}