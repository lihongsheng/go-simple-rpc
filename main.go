package main

import (
	"fmt"
	client2 "go-simple-rpc/client"
	"go-simple-rpc/common"
	"go-simple-rpc/service"
	"net"
	"sync"
	"time"
)

func main() {
	// server 启动
	service.AddReadHandler(common.Decode) //接码
	// 限流在服务端，接受请求的地方
	// 熔断在调用端，发起调用的地方
	// 降级可以在配置端下放配置到各个服务，也可以在调用端【熔断处】上报
	// 所以需要在这里加限流的处理，未加入此模块
	// 调用端需要做连接池，保证TCP是单向的。否则需要根据每次的请求ID做队列等待
	service.AddReadHandler(service.RpcHandler) //
	service.AddWriteHandler(common.Encode)
	wg := sync.WaitGroup{}
	wg.Add(2)
	service.RegisterServer("user",new(service.UserServer))

	go func() {
		defer wg.Done()
		service.DefaultServer.Listen(9099,"tcp")
		service.DefaultServer.Run()
	}()
	go func() {
		defer wg.Done()
		// 调用端
		time.Sleep(5 * time.Second)
		con,_ := net.Dial("tcp","127.0.0.1:9099")
		userClient := client2.NewUser(con)
		fmt.Println(userClient.GetUser(1))
	}()

	wg.Wait()
}
