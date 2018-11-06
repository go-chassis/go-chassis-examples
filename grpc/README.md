# Quick start
this example shows how grpc-go leverage go chassis to become a cloud native application

1. Launch service center

2. Run server

```sh 
cd examples/grpc/server
export CHASSIS_HOME=$PWD
go run main.go

```

3. Run client
```sh 
 cd examples/grpc/client
 export CHASSIS_HOME=$PWD
 go run main.go
 
```