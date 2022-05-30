package main

import (
	"context"
	"github.com/go-chassis/openlog"
	"net/http"

	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulRouterB is a struct used for implementation of restfull router program
type RestFulRouterB struct {
	called int
}

// Equal is method to compare given num and slice product
func (r *RestFulRouterB) GetPayments(ctx *restful.Context) {
	req, err := rest.NewRequest("GET", "http://paymentService/payments", nil)
	if err != nil {
		openlog.Error("new request failed.")
		return
	}

	resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
	if err != nil {
		openlog.Error("do request failed.")
		return
	}
	defer resp.Body.Close()

	ctx.Write(httputil.ReadBody(resp))
}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterB) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/payments", ResourceFunc: r.GetPayments},
	}
}

func main() {
	chassis.RegisterSchema("rest", &RestFulRouterB{})
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
