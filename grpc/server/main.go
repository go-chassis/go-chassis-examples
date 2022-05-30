package main

import (
	pb "github.com/go-chassis/go-chassis-examples/grpc/helloworld"
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/openlog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/grpc/server/
// Server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md["x-user"])
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func main() {
	chassis.RegisterSchema("grpc", &Server{}, server.WithRPCServiceDesc(&pb.Greeter_serviceDesc))
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed.")
		return
	}
	chassis.Run()
}
