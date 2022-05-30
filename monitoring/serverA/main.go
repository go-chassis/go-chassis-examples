package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/go-chassis/v2/examples/schemas"
	"github.com/go-chassis/openlog"

	//tracers
	_ "github.com/go-chassis/go-chassis-extension/tracing/zipkin"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/serverA/

func main() {
	chassis.RegisterSchema("rest", &schemas.TracingHello{}, server.WithSchemaID("TracingHello"))
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
