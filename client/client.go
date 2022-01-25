package client

import (
	"encoding/binary"
	"encoding/json"
	"go-simple-rpc/common"
	"go-simple-rpc/common/protocol"
	"go-simple-rpc/common/snowflake"
	"net"
)

type UserClientProxy struct {
	con net.Conn
}

func NewUser(con net.Conn) UserClientProxy  {
	return UserClientProxy{
		con: con,
	}
}

func (u UserClientProxy) GetUser(id int) common.User {
	p := protocol.Protocol{}
	p.SetVersion(protocol.Version)
	rid := snowflake.GenerateId(1,1,false)
	p.SetRequestId(rid)
	p.SetBodyType(1)
	p.SetRequestType(1)
	pa := make([]interface{},1)
	pa[0] = id
	re := common.Request{
		ServerName: "user",
		MethodName: "GetUser",
		Params: pa,
	}
	body,_ := json.Marshal(re)
	p.SetBody(body)
	netBody := p.SerializeToByte()
	u.write(netBody)

	response := u.read()
	return response.Params[0].(common.User)
}

func (u UserClientProxy)  write(message []byte)  {
	u.con.Write(message)
}

func (u UserClientProxy) read() (r common.Response) {
	h := make([]byte,18)
	//var err error
	_,err := u.con.Read(h)
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
	_,err = u.con.Read(message)
	if err != nil {
		return
	}
	p.SetBody(message)

	json.Unmarshal(message,&r)
	return r
}
