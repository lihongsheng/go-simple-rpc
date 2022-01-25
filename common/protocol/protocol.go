package protocol

import (
	"encoding/binary"
	"fmt"
)

// 一个完整的报文
// 版本号 + 标识长度字节 + 请求ID + 请求类型 + 报文编码类型 + 实际报文
const (
	//协议头
	HEAD_SIZE int = 4
	//协议包的所在长度值为 [ 请求ID的8个字节 +请求type的一个字节 +  报文的编码类型  + body的字节数]
	PROTOCOL_LENGTH = 4
	//请求ID的长度
	REQUESID_LENGTH = 8
	//请求的类型的长度
	REQUEST_TYPE = 1
	// 报文的编码类型 标识是json,protobuf,god等等编码
	BODY_TYPE = 1
)
const Version = 0xabef0101
//

type Protocol struct {
	message []byte
	version int
	requestId int64
	protocolLength int
	requestType int8
	bodyType int8
}

func (p Protocol) GetVersion() int {
	return p.version
}

func (p *Protocol) SetVersion(ver int) {
	p.version = ver
}

func (p *Protocol) SetRequestId(id int64)  {
	p.requestId = id
	//binary.BigEndian.PutUint64(p.message[8:12],uint64(Version))
}

func (p *Protocol) GetRequestId() int64   {
	return p.requestId
}

func (p *Protocol) GetRequestType() int8   {
	return p.requestType
}

func (p *Protocol) SetRequestType(id int )    {
	p.requestType = int8(id)
}


func (p *Protocol) GetBodyType() int8   {
	return p.bodyType
}

func (p *Protocol) SetBodyType(id int )    {
	p.bodyType = int8(id)
}

func (p *Protocol) SetBody(body []byte)  {
	var length int
	length = len(body) + REQUESID_LENGTH + REQUEST_TYPE + BODY_TYPE
	p.message = body
	p.protocolLength = length
}

func (p Protocol) GetBody() []byte {
	return p.message
}

func (p *Protocol) SerializeToByte() []byte {
	body := make([]byte,18)
	fmt.Println(8+p.protocolLength)
	// 写入头部
	binary.BigEndian.PutUint32(body[0:4],uint32(p.version))
	binary.BigEndian.PutUint32(body[4:8],uint32(p.protocolLength))
	binary.BigEndian.PutUint64(body[8:16],uint64(p.requestId))
	body[16] = byte(p.requestType)
	body[17] = byte(p.bodyType)
	// 设置报文
	body = append(body,p.message...)
	return body
}
//func (p Protocol) SetVersion() int64 {
//
//}
