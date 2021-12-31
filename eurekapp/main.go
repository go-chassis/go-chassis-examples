package main

import (
	"eurekapp/controller"
	"eurekapp/router"

	"gitee.com/daqingshu/eurekaplugin/v2"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/openlog"
)

func main() {
	eurekaplugin.Init()
	chassis.RegisterSchema("rest", &controller.Data{}, server.WithSchemaID("PriceService"))
	chassis.RegisterSchema("rest", &router.RestFulRouterA{}, server.WithSchemaID("RestROUTERService"))
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
