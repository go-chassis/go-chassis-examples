this demo shows how to leverage router management to make a canary release

we have 2 version of cart service
in folders serviceV1 and service V2

the response is different 
```go
// Equal is method to compare given num and slice sum
func (r *RestFulRouterA) Get(context *restful.Context) {
	context.Write([]byte("version 1.0"))
}

```
```go
// Equal is method to compare given num and slice sum
func (r *RestFulRouterA) Get(context *restful.Context) {
	context.Write([]byte("version 1.1"))
}

```

here is a configuration is consumer side

```yaml
routeRule:
  carts:
    - precedence: 2
      match:
        headers:
          Os:
            regex: ios
      route:
        - weight: 100
          tags:
            version: 1.1
    - precedence: 1
      route:
        - weight: 100
          tags:
            version: 1.0
```

it means if header key "Os"'s  value is ios
then all the request should be route to version 1.1.
Other request router to 1.0

you can build and run to see results

```shell
cd serverV1
go build main.go
./main
```
```shell
cd serverV2
go build main.go
./main
```

```shell
cd client
go build main.go
./main
```