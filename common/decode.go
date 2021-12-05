package common

import (
	"go-simple-rpc/common/protocol"
	"io"
)

func Decode( reader io.ReadWriteCloser) *protocol.Protocol {

	version := reader.Read()
}