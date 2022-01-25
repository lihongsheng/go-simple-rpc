package service

import (
	"encoding/json"
	"fmt"
	"go-simple-rpc/common"
	"go-simple-rpc/common/protocol"
	"reflect"
)


func RpcHandler(ct *common.LinkHandler,param interface{})  {
	p := param.(*protocol.Protocol)
	r := &common.Request{}
	json.Unmarshal(p.GetBody(),r) // 接码字节，反解码 request对象
	fmt.Println(r.ServerName)
	ser := ServerList[r.ServerName]

	vOf := reflect.ValueOf(ser)
	// 构造参数
	funcParam := make([]reflect.Value,0)
	for _,pa := range r.Params {
		tmp := int(pa.(float64))
		funcParam = append(funcParam,reflect.ValueOf(tmp))
	}
	ret := vOf.MethodByName(r.MethodName).Call(funcParam)
	realReturn := make([]interface{},0)
	// 结果集转换
	for _,rr := range ret {
		realReturn = append(realReturn,rr.Interface().(common.User))
	}
	response := common.Response{
		Params: realReturn,
	}
	res,_ := json.Marshal(response)
	p.SetBody(res)
	ct.Write(p)
	// 基于serverName找寻到server
	// 如果server有调用其他服务的server需要加入熔断机制
	// 反射调用server的方法并传入参数
	// 执行结果生成 response 结构体
	// 返回数据，并再次序列化
}