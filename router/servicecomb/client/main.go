package main

import (
	"context"
	"github.com/go-chassis/go-chassis/server/restful"
	"net/http"

	//_ "github.com/apache/servicecomb-kie/client/adaptor"

	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
)

// RestFulRouterB is a struct used for implementation of restfull router program
type RestFulRouterB struct {
	called int
}

// Equal is method to compare given num and slice product
func (r *RestFulRouterB) GetPayments(ctx *restful.Context) {
	req, err := rest.NewRequest("GET", "http://paymentService/payments", nil)
	if err != nil {
		lager.Logger.Error("new request failed.")
		return
	}

	resp, err := core.NewRestInvoker().ContextDo(context.Background(), req)
	if err != nil {
		lager.Logger.Error("do request failed.")
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
	chassis.RegisterSchema("rest",&RestFulRouterB{})
	//Init framework
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
