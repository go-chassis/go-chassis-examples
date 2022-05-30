package resource

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

// RestFulRouterA is a struct used for implementation of restfull router program
type RestFulRouterB struct {
}

// Equal is method to compare given num and slice sum
func (r *RestFulRouterB) Get(context *restful.Context) {
	context.Write([]byte("version 1.1"))
}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterB) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/list", ResourceFuncName: "Get"},
	}
}
