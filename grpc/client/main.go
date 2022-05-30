package main

import (
	"github.com/go-chassis/go-chassis-examples/grpc/helloworld"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/client"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/core/common"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/grpc/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	//declare reply struct
	reply := &helloworld.HelloReply{}
	//header will transport to target service
	ctx := common.NewContext(map[string]string{
		"X-User": "pete",
	})
	invoker := core.NewRPCInvoker()

	//Invoke with micro service name, schema ID and operation ID
	if err := invoker.Invoke(ctx, "RPCServer", "helloworld.Greeter", "SayHello",
		&helloworld.HelloRequest{Name: "Peter"}, reply, core.WithProtocol("grpc")); err != nil {
		openlog.Error("error" + err.Error())
	}
	openlog.Info(reply.Message)
}
