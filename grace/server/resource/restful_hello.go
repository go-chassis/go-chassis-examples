package resource

import (
	"net/http"

	rf "github.com/go-chassis/go-chassis/server/restful"
	"time"
)

//GraceResource is a struct used for implementation of restful api
type GraceResource struct {
}

func (r *GraceResource) Wait(b *rf.Context) {
	time.Sleep(10 * time.Second)
	b.Write([]byte("success"))
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *GraceResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/grace", ResourceFuncName: "Wait"},
	}
}
