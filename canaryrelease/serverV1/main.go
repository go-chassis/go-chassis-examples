package main

import (
	"github.com/go-chassis/go-chassis-examples/canaryrelease/serverV1/resource"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rpc/server/

func main() {
	chassis.RegisterSchema("rest", &resource.RestFulRouterA{})
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
