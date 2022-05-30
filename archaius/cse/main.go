package main

import _ "github.com/huaweicse/auth/adaptor/gochassis"
import (
	"github.com/go-chassis/go-chassis-examples/archaius/resource"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/openlog"
)

func main() {
	chassis.RegisterSchema("rest", &resource.Hello{}, server.WithSchemaID("Hello"))
	err := chassis.Init()
	if err != nil {
		openlog.Error(err.Error())
		return
	}
	chassis.Run()
}
