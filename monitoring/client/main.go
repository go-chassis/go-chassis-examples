package main

import (
	"context"
	"github.com/go-chassis/openlog"

	_ "github.com/go-chassis/go-chassis-extension/tracing/zipkin"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"time"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}

	req, err := rest.NewRequest("GET", "http://RESTServerA/trace", nil)
	if err != nil {
		openlog.Error("new request failed." + err.Error())
		return
	}

	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		openlog.Error("do request failed." + err.Error())
		return
	}
	defer resp.Body.Close()
	openlog.Info("REST Server sayhello[GET]: " + string(httputil.ReadBody(resp)))
	time.Sleep(2 * time.Second)
}
