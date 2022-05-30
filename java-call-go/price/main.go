package main

import (
	"github.com/go-chassis/go-chassis-examples/java-call-go/price/controller"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/lager"
	"github.com/go-chassis/go-chassis/v2/core/server"
)

func main() {
	chassis.RegisterSchema("rest", &controller.Data{}, server.WithSchemaID("PriceService"))
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
