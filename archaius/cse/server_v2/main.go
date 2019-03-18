package main

import _ "github.com/huaweicse/auth/adaptor/gochassis"
import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/archaius/resource"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/
func main() {
	chassis.RegisterSchema("rest", &resource.Hello{})
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
