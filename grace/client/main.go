package main

import (
	"context"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}

	req, err := rest.NewRequest("GET", "http://GraceService/grace", nil)
	if err != nil {
		openlog.Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			openlog.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
			defer resp.Body.Close()
		}
	}

	if err != nil {
		openlog.GetLogger().Error("request failed: " + err.Error())
		return
	}
}
