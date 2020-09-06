package main

import (
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
)

type HelloResource struct {
}

func (r *HelloResource) SayHi(b *rf.Context) {
	b.Write([]byte("hello. go chassis"))
	return
}

func (r *HelloResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hello", ResourceFunc: r.SayHi},
	}
}

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/

func main() {
	chassis.RegisterSchema("rest", &HelloResource{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
