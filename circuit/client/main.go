package main

import (
	"github.com/go-chassis/go-chassis"

	"context"

	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-mesh/openlogging"
	"sync"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/circuit/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}
	for i := 0; i < 100; i++ {
		callError()
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go callConcurrency(wg)
	}
	wg.Wait()
	for i := 0; i < 20; i++ {
		callDeadlock()
	}

}
func callError() {
	req, err := rest.NewRequest("GET", "http://CircuitServer/error", nil)
	if err != nil {
		openlogging.GetLogger().Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
			defer resp.Body.Close()
		}
	}

	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return
	}

}
func callConcurrency(wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := rest.NewRequest("GET", "http://CircuitServer/concurrency", nil)
	if err != nil {
		openlogging.Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
			openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
		}
	}
	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return
	}

}
func callDeadlock() {
	req, err := rest.NewRequest("GET", "http://CircuitServer/lock", nil)
	if err != nil {
		openlogging.Error("new request failed.")
		return
	}
	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if resp != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
			openlogging.GetLogger().Info("response body: " + string(httputil.ReadBody(resp)))
		}
	}
	if err != nil {
		openlogging.GetLogger().Error("request failed: " + err.Error())
		return
	}

}
