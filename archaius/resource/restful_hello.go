package resource

import (
	"net/http"

	"fmt"
	"strconv"
	"sync"
	"time"

	"errors"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	rf "github.com/go-chassis/go-chassis/server/restful"
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
func (r *Hello) Delay(b *rf.Context) {
	// test circuit config
	delay := b.ReadPathParameter("delay")
	sleep, err := strconv.Atoi(delay)
	if err != nil {
		b.WriteError(http.StatusBadRequest, err)
		return
	}
	sleepTime := time.Duration(sleep) * time.Millisecond
	time.Sleep(sleepTime)
	b.WriteHeader(http.StatusOK)
	b.Write([]byte(fmt.Sprintf("instance :%v ,version: %v , sleep %v , reply success", runtime.InstanceID,runtime.Version, sleepTime)))
}
func (r *Hello) Error(b *rf.Context) {
	//b.WriteHeader()
	b.WriteError(http.StatusInternalServerError, errors.New(fmt.Sprintf("instance :%v ,error", runtime.InstanceID)))
}

//URLPatterns helps to respond for corresponding API calls
func (r *Hello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/file", ResourceFuncName: "GetFileContent"},
		{Method: http.MethodGet, Path: "/props/{key}", ResourceFuncName: "GetProps"},
		{Method: http.MethodGet, Path: "/delay/{delay}", ResourceFuncName: "Delay"},
		{Method: http.MethodGet, Path: "/error", ResourceFuncName: "Error"},
	}
}
