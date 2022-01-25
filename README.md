# go-simple-rpc
一个从零构建的简单的go-rpc示例

### 协议的格式
[版本号]|[数据包长度]|[requestID]|[请求报文类型]|[报文编码格式]|[实际报文内容]的封包格式
### request 请求体结构
```go
type Request struct {
	ServerName string
	MethodName string
	Params []interface{}
}
```
### 一个RPC的请求流程
客户端建立与服务端的链接 -> 基于客户端代理类，调用需要请求的方法 -> 生成请求体request，基于编码格式和协议格式，做数据转换->发送到服务端-》服务端接受到二进制报文，基于编码和协议格式，做反解码操作-》基于request的指定的server类和方法名调用注册的server方法返回数据-> 服务端把返回数据，编码返回给客户端->客户端反解码，拿到调用结果
### 服务端

### 客户端
