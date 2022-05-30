package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"github.com/go-chassis/go-chassis-examples/grpc/helloworld"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/client"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/openlog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var wg = sync.WaitGroup{}

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/grpc/client/
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	n := 1000000
	t := 10
	times := n / t

	for i := 0; i < t; i++ {
		wg.Add(1)
		go call(times)
	}
	wg.Wait()
}

func call(times int) {
	invoker := core.NewRPCInvoker()
	//declare reply struct
	for i := 0; i < times; i++ {
		reply := &helloworld.HelloReply{}
		ctx := context.Background()
		//Invoke with micro service name, schema ID and operation ID
		if err := invoker.Invoke(ctx, "RPCServer",
			"helloworld.Greeter", "SayHello",
			&helloworld.HelloRequest{Name: "Peter"}, reply,
			core.WithProtocol("grpc")); err != nil {
			openlog.Error("error" + err.Error())
		}
	}
	//wg.Done()
}
func call2(times int) {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "peter"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//declare reply struct
	for i := 0; i < times; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		_, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	//wg.Done()
}
