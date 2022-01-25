package common

import "go-simple-rpc/common/protocol"

// 编码
func Encode(ct *LinkHandler,param interface{})  {
	p := param.(*protocol.Protocol)
	// 写入数据
	ct.Con.Write(p.SerializeToByte())
}
