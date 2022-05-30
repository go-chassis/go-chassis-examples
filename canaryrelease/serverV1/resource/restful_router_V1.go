package resource

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

// RestFulRouterA is a struct used for implementation of restfull router program
type RestFulRouterA struct {
}

// Equal is method to compare given num and slice sum
func (r *RestFulRouterA) Get(context *restful.Context) {
	context.Write([]byte("version 1.0"))
}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterA) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/list", ResourceFuncName: "Get"},
	}
}
