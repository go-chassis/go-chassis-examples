Graceful shutdown http server

1.Launch service center

follow https://github.com/apache/servicecomb-service-center/tree/master/examples/infrastructures/docker

2.Run server

```sh 
cd server
export CHASSIS_HOME=$PWD
go run main.go

```

3.Run client
```sh 
cd client
export CHASSIS_HOME=$PWD
go run main.go
```

server will wait 10 seconds to return a string "success", if you kill server before it response, 
the server will wait until all requests is finished.
you will see success log 

INFO client/main.go:29 response body: success

