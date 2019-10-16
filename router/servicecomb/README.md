# Router example

In this example we will run two different versions service and one client.We will show you how to 
implement grayscale publishing of the application with route management,version V1 is simulation
you old service , version V2 is simulation  you new service. 
1. Build and Run Service V1 and Service V2
 
 
 we use microservice.yaml to set service config,and we must set two service with the 
 same service name,set different version for two service
 
 
Service V1
```yaml
#Private property of microservices
service_description:
  name: ROUTERServer
  version: 1.0
```
 
Service V2
```yaml
#Private property of microservices
service_description:
  name: ROUTERServer
  version: 2.0
```
 
 build and run two service
 
```bash
cd serverV1
go build main.go
./main 
```
```bash
cd serverV2
go build main.go
./main 
```

2. Build and run client

client use router.yaml to management router.Configuration of this file 
use router.yaml launch client 
```yaml
servicecomb:
  routeRule:
    paymentService: |
      - precedence: 1 # 优先级，数字越大优先级越高
        route: #路由规则列表
        - tags:
            version: 1.0
            project: x
          weight: 50 #全重 50%到这里
        - tags:
            version: 2.0
            project: z
          weight: 50 #全重 50%到这里
```

build and run client
```bash
cd client
go build main.go
./main 
```


result:

{"level":"INFO","timestamp":"2019-10-16 12:57:29.893 +08:00","file":"client/main.go:43","msg":"paymentService response: version V1 was called: 1 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:30.896 +08:00","file":"client/main.go:43","msg":"paymentService response: version V2 was called: 1 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:31.899 +08:00","file":"client/main.go:43","msg":"paymentService response: version V1 was called: 2 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:32.901 +08:00","file":"client/main.go:43","msg":"paymentService response: version V2 was called: 2 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:33.903 +08:00","file":"client/main.go:43","msg":"paymentService response: version V1 was called: 3 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:34.905 +08:00","file":"client/main.go:43","msg":"paymentService response: version V2 was called: 3 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:35.906 +08:00","file":"client/main.go:43","msg":"paymentService response: version V1 was called: 4 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:36.908 +08:00","file":"client/main.go:43","msg":"paymentService response: version V2 was called: 4 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:37.910 +08:00","file":"client/main.go:43","msg":"paymentService response: version V1 was called: 5 times"}
{"level":"INFO","timestamp":"2019-10-16 12:57:38.910 +08:00","file":"client/main.go:43","msg":"paymentService response: version V2 was called: 5 times"}
