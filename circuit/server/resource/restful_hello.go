package resource

import (
	"errors"
	"net/http"

	rf "github.com/go-chassis/go-chassis/server/restful"
	"sync"
)

//CircuitResource is a struct used for implementation of restful api
type CircuitResource struct {
	lock sync.Mutex
}

//Sayerror is a method used to reply request user with error, it fulfills client side error threshold
func (r *CircuitResource) GetError(b *rf.Context) {
	b.WriteError(http.StatusInternalServerError, errors.New("internal error"))
	return
}

//GetLock cause dead lock, it will cause client side timeout
func (r *CircuitResource) GetLock(b *rf.Context) {
	r.lock.Lock()
	return
}

//MaxConcurrency is for MaxConcurrency test
func (r *CircuitResource) MaxConcurrency(b *rf.Context) {
	b.WriteHeader(http.StatusOK)
	b.Write([]byte("too much concurrency"))
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *CircuitResource) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/error", ResourceFuncName: "GetError"},
		{Method: http.MethodGet, Path: "/lock", ResourceFuncName: "GetLock"},
		{Method: http.MethodGet, Path: "/concurrency", ResourceFuncName: "MaxConcurrency"},
	}
}
