package main

import (
	"context"
	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}

	req, err := rest.NewRequest("GET", "http://GraceService/grace", nil)
	if err != nil {
		openlogging.Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
			defer resp.Body.Close()
		}
	}

	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return
	}
}
