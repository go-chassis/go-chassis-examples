package resource

import "net/http"
import (
	"context"

	"sync"

	"fmt"
	"time"

	"strconv"

	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	rf "github.com/go-chassis/go-chassis/server/restful"
		)

const (
	Router      = "router"
	Circuit     = "circuit"
	Rate        = "rate"
	Fault       = "fault"
	LoadBalance = "load_balance"
)

type HelloClient struct {
	lock sync.Mutex
}
type result struct {
	Reply string
	Error string
	Num   int
}
type Reply struct {
	TestType string    `json:"testType"`
	Reply    *[]result `json:"reply"`
}

var lock sync.Mutex
var restInvoker = core.NewRestInvoker()

func (client *HelloClient) TestRouter(b *rf.Context) {
	testFunc(b, Router)
}
func (client *HelloClient) TestCircuit(b *rf.Context) {
	testFunc(b, Circuit)

}
func (client *HelloClient) TestRate(b *rf.Context) {
	testFunc(b, Rate)

}
func (client *HelloClient) TestFault(b *rf.Context) {
	testFunc(b, Fault)
}
func (client *HelloClient) TestLB(b *rf.Context) {
	testFunc(b, LoadBalance)
}

func testFunc(b *rf.Context, tp string) {
	// get access time like 1s,10ms,
	duration, err := time.ParseDuration(b.ReadQueryParameter("d"))
	if err != nil {
		panic(err)
	}
	// get goroutine count
	count, err := strconv.Atoi(b.ReadQueryParameter("c"))
	if err != nil {
		panic(err)
	}
	// get delay time
	delay := b.ReadQueryParameter("delay")

	cancels := make([]context.CancelFunc, 0)
	results := &[]result{}

	url := fmt.Sprintf("http://RESTServer/delay/%s", delay)
	for i := 0; i < count; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		cancels = append(cancels, cancel)
		go callHTTP(ctx, restInvoker, http.MethodGet, url, tp, results)
	}
	// locked before time
	t := time.Tick(duration)
	<-t
	for _, c := range cancels {
		c()
	}
	replys := Reply{
		TestType: tp,
		Reply:    results,
	}
	b.WriteJSON(replys, "application/json")
}

func callHTTP(ctx context.Context, restInvoker *core.RestInvoker, method string, url string, t string, r *[]result) {
	for {

		req, err := rest.NewRequest(method, url, nil)
		if err != nil {
			panic(err)
		}
		if t == Router {
			req.Header.Set("Foo", "bar")
		}

		resp, err := restInvoker.ContextDo(ctx, req)

		if err != nil {
			appendReult(err.Error(), "", r)
			if !endOfCtx(ctx) {
				break
			}
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode > 300 {
			appendReult(string(httputil.ReadBody(resp)), "", r)
			if !endOfCtx(ctx) {
				break
			}
			continue
		}
		appendReult("",fmt.Sprintf(string(httputil.ReadBody(resp))+"time :%d", time.Now().UnixNano()/1e6),r)
		resp.Body.Close()
		if !endOfCtx(ctx) {
			break
		}
	}
}
func appendReult(errMsg, replyMsg string, r *[]result) {
	lock.Lock()
	defer lock.Unlock()
	if errMsg != "" {
		*r = append(*r, result{Error: errMsg, Num: len(*r) + 1})
		return
	}
	*r = append(*r, result{Reply: replyMsg, Num: len(*r) + 1})
}
func endOfCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		return true
	}
}

//URLPatterns helps to respond for corresponding API calls
func (client *HelloClient) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/client/router", ResourceFuncName: "TestRouter"},
		{Method: http.MethodGet, Path: "/client/circuit", ResourceFuncName: "TestCircuit"},
		{Method: http.MethodGet, Path: "/client/rate", ResourceFuncName: "TestRate"},
		{Method: http.MethodGet, Path: "/client/fault", ResourceFuncName: "TestFault"},
		{Method: http.MethodGet, Path: "/client/load_balance", ResourceFuncName: "TestLB"},
	}
}