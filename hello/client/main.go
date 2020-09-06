package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type SimpleResource struct {
}

func (r *SimpleResource) SayHi(b *rf.Context) {
	req, _ := rest.NewRequest(http.MethodGet, "http://HelloServer/hello", nil)
	restInvoker := core.NewRestInvoker()
	resp, err := restInvoker.ContextDo(context.TODO(), req)
	if err != nil {
		log.Println(err)
		return
	}
	b.Write(httputil.ReadBody(resp))
	return
}

func (r *SimpleResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hi", ResourceFunc: r.SayHi},
	}
}

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/

func main() {
	chassis.RegisterSchema("rest", &SimpleResource{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
