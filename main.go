package main

import (
	"flag"
	"fmt"
	client2 "go-simple-rpc/client"
	"go-simple-rpc/common"
	"go-simple-rpc/service"
	"net"
)

// 服务端启动
func startServer(port int)  {
	// server 启动
	service.AddReadHandler(common.Decode) //接码
	// 限流在服务端，接受请求的地方
	// 熔断在调用端，发起调用的地方
	// 降级可以在配置端下放配置到各个服务，也可以在调用端【熔断处】上报
	// 所以需要在这里加限流的处理，未加入此模块
	// 调用端需要做连接池，保证TCP是单向的。否则需要根据每次的请求ID做队列等待
	service.AddReadHandler(service.RpcHandler) // 注册读取客户端数据处理函数
	service.AddWriteHandler(common.Encode) // 注册回写客户端数据的处理函数
	service.RegisterServer("user",new(service.UserServer)) // 微服务
	service.DefaultServer.Listen(port,"tcp")
	service.DefaultServer.Run()
}

// 客户端启动
func startClient(port int) {
	con,err := net.Dial("tcp",fmt.Sprintf("127.0.0.1:%d",port))
	if err == nil {
		fmt.Println("err: ",err.Error())
	}
	fmt.Println("client success")
	userClient := client2.NewUser(con)
	// 打印值
	fmt.Println(userClient.GetUser(1))
}
var start = flag.String("start","client","请输入client或者server")
var port = flag.Int("port",9099,"请输入端口号")

func main() {
	flag.Parse() //接受和解析 命令行参数
}
