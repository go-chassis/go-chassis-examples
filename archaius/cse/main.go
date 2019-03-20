package main

import _ "github.com/huaweicse/auth/adaptor/gochassis"
import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/archaius/resource"
	"github.com/go-chassis/go-chassis/core/server"
	"github.com/go-mesh/openlogging"
)

func main() {
	chassis.RegisterSchema("rest", &resource.Hello{}, server.WithSchemaID("Hello"))
	err := chassis.Init()
	if err != nil {
		openlogging.GetLogger().Error(err.Error())
		return
	}
	chassis.Run()
}
