package router

import (
	"net/http"

	"github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulRouterB is a struct used for implementation of restfull router program
type RestFulRouterA struct {
}

// Equal is method to compare given num and slice product
func (r *RestFulRouterA) Info(context *restful.Context) {
	versionInfo := struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		HostName string `json:"hostName"`
	}{
		Name:     "CHASSIS_SERVER_V1",
		Version:  "1.0",
		HostName: context.ReadRequest().Host,
	}

	context.WriteJSON(versionInfo, "application/json")

}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterA) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/info", ResourceFunc: r.Info},
	}
}
