package schema

import (
	"net/http"

	"github.com/go-chassis/go-chassis/server/restful"
)

type Server struct{}

func (*Server) PathParm(ctx *restful.Context) {
	name := ctx.ReadPathParameter("name")
	ctx.Write([]byte("name for path param is : " + name))
}

func (*Server) DynamicURL(ctx *restful.Context) {
	ctx.Write([]byte("dynamic path : " + ctx.ReadPathParameter("dynamic")))
}

func (*Server) DynamicURLLetter(ctx *restful.Context) {
	ctx.Write([]byte("dynamic path : " + ctx.ReadPathParameter("dynamic")))
}

func (*Server) AllAccess(ctx *restful.Context) {
	ctx.Write([]byte("dynamic path : " + ctx.ReadPathParameter("dynamic")))
}
func (*Server) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodGet, Path: "/param/{name}", ResourceFuncName: "PathParm"},
		{Method: http.MethodGet, Path: "/dynamic/{dynamic:*}", ResourceFuncName: "DynamicURL"},
		{Method: http.MethodGet, Path: "/dynamic_letter/{dynamic:[a-zA-Z]}", ResourceFuncName: "DynamicURLLetter"},
		{Method: http.MethodGet, Path: "/{dynamic:*}", ResourceFuncName: "AllAccess"},
	}
}
