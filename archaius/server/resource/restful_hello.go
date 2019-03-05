package resource

import (
	"net/http"

	"github.com/go-chassis/go-archaius"
	rf "github.com/go-chassis/go-chassis/server/restful"
	"sync"
)

type Hello struct {
	lock sync.Mutex
}

func (r *Hello) GetFileContent(b *rf.Context) {
	b.WriteHeader(http.StatusOK)
	//get content of a file
	b.Write([]byte(archaius.GetString("custom.yaml", "default")))
	return
}
func (r *Hello) GetProps(b *rf.Context) {
	b.WriteHeader(http.StatusOK)
	k := b.ReadPathParameter("key")
	//get single key value
	b.Write([]byte(archaius.GetString(k, "default")))
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *Hello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/file", ResourceFuncName: "GetFileContent"},
		{Method: http.MethodGet, Path: "/props/{key}", ResourceFuncName: "GetProps"},
	}
}
