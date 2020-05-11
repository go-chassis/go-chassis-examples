package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis-examples/metrics/server/schema"
	"github.com/go-chassis/go-chassis/core/server"
	"github.com/go-chassis/go-chassis/pkg/metrics"
	"github.com/go-mesh/openlogging"
)

func main() {
	chassis.RegisterSchema("rest", &schema.User{}, server.WithSchemaID("server"))
	err := chassis.Init()
	if err != nil {
		openlogging.GetLogger().Errorf("chassis init failed ,err :%v", err.Error())
	}
	metrics.CreateCounter(metrics.CounterOpts{
		Help:   "count user login",
		Name:   schema.Login,
		Labels: []string{schema.Label}})
	metrics.CreateCounter(metrics.CounterOpts{
		Help:   "user sign out",
		Name:   schema.SignOut,
		Labels: []string{schema.Label}})
	chassis.Run()
}
