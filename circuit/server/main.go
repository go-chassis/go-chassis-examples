package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/circuit/server/resource"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/circuit/server/
func main() {
	chassis.RegisterSchema("rest", &resource.CircuitResource{})
	//start all server you register in server/schemas.
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
