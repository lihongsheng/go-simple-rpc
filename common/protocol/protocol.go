package protocol

import "encoding/binary"

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
}

func (p Protocol) GetVersion() int64 {
	return Version
}

func (p *Protocol) SetVersion() {
	if p.message == nil {
		p.message = make([]byte,18)
	}
	binary.BigEndian.PutUint32(p.message[0:4],uint32(Version))
}

func (p *Protocol) SetRequestId(id int64)  {
	binary.BigEndian.PutUint64(p.message[8:12],uint64(Version))
}

func (p *Protocol) GetRequestId() int64   {
	id := binary.BigEndian.Uint64(p.message[8:12])
	return int64(id)
}

func (p *Protocol) GetRequestType() int   {
	id := int(p.message[17])
	return id
}

func (p *Protocol) SetRequestType(id int )    {
	p.message[17] = byte(id)
}


func (p *Protocol) GetBodyType() int   {
	id := p.message[18]
	return int(id)
}

func (p *Protocol) SetBodyType(id int )    {
	p.message[18] = byte(id)
}

func (p *Protocol) SetBody(body []byte)  {
	var length int
	length = len(body) + REQUESID_LENGTH + REQUEST_TYPE + BODY_TYPE
	binary.BigEndian.PutUint32(p.message[4:8],uint32(length))
}

func (p Protocol) GetBody() []byte {
	return p.message[18:]
}

func (p Protocol) GetDecodeBody() interface{} {
	//return p.message[18:]
	// 基于编码类型进行接码
	return nil
}
//func (p Protocol) SetVersion() int64 {
//
//}
