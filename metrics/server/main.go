package main

import (
	"github.com/go-chassis/go-chassis-examples/metrics/server/schema"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/go-chassis/v2/pkg/metrics"
	"github.com/go-chassis/openlog"
)

func main() {
	chassis.RegisterSchema("rest", &schema.User{}, server.WithSchemaID("server"))
	err := chassis.Init()
	if err != nil {
		openlog.Error("chassis init failed ,err: " + err.Error())
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
