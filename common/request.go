package common

type Request struct {
	ServerName string
	MethodName string
	Params []interface{}
}
