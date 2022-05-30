package main

import (
	"context"
	"time"

	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/client/

// Implement grayscale publishing of the application,version  A is you old service ,version B is you
// new service.you want to small request to access you new service to test new service of new function

func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		req, err := rest.NewRequest("GET", "http://carts/list", nil)
		if err != nil {
			openlog.Error("new request failed.")
			return
		}

		//req.Header.Set("Chassis", "info")

		resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
		if err != nil {
			openlog.Error("do request failed.")
			return
		}
		defer resp.Body.Close()
		openlog.Info("carts Server return : " + string(httputil.ReadBody(resp)))

		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 10; i++ {
		req, err := rest.NewRequest("GET", "http://carts/list", nil)
		if err != nil {
			openlog.Error("new request failed.")
			return
		}

		req.Header.Set("os", "ios")

		resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
		if err != nil {
			openlog.Error("do request failed.")
			return
		}
		defer resp.Body.Close()
		openlog.Info("carts Server return : " + string(httputil.ReadBody(resp)))

		time.Sleep(1 * time.Second)
	}
}
