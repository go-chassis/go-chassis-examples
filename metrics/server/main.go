package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/metrics/server/schema"
	"github.com/go-chassis/go-chassis/core/server"
	"github.com/go-mesh/openlogging"
)

func main() {
	chassis.RegisterSchema("rest", &schema.User{}, server.WithSchemaID("server"))
	err := chassis.Init()
	if err != nil {
		openlogging.GetLogger().Errorf("chassis init failed ,err :%v", err.Error())
	}
	chassis.Run()
}
