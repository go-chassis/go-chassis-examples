package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/grpc/helloworld"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	_ "github.com/go-chassis/go-chassis/client/grpc"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/common"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/grpc/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
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
		openlogging.Error("error" + err.Error())
	}
	openlogging.Info(reply.Message)
}
