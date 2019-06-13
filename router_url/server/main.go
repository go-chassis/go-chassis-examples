package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/router_url/server/schema"
	"github.com/go-chassis/go-chassis/core/server"
)

func main() {
	chassis.RegisterSchema("rest", &schema.Server{}, server.WithSchemaID("test"))
	if err := chassis.Init(); err != nil {
		panic(err)
	}
	chassis.Run()
}
