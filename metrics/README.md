# Quik start with examples

1. Launch service center
```sh
cd examples
docker-compose up
```

2. Run rest server

```sh 
cd server
export CHASSIS_HOME=$PWD
go run main.go

```

3. call api `/login` or `sign_out`
api of `/login` is `POST` , `sign_out` is `GET`

call `http://127.0.0.1:8080/login` , the data of body 
```json
{
	"user_name":"user_test",
	"sex":"male",
	"age":27,
	"pwd":"123456"
}
```
will reply : 
```json
{
    "user_name": "user_test",
    "sex": "male",
    "age": 27
}
``` 
4. get metrics data
if you successful access the api `/login` , then you access `http://127.0.0.1:8080/metrics ` ,
 in the return content you can see the following 
```text
# HELP login_total count user login
# TYPE login_total counter
login_total{username="user_test login"} 1
```
5. [metrics](https://docs.go-chassis.com/user-guides/metrics.html) API

before you use metrics api of go-chassis , you need use `metrics.Init()` init metrics

##### create counter

```go
metrics.CreateCounter(metrics.CounterOpts{
    Help:   "count user login",
    Name:   "login_total",
    Labels: []string{"username"}})
```

add data into metrics
```go
metrics.CounterAdd("sign_out_total", 1, map[string]string{"username": fmt.Sprintf("%s sign out", "user_test")})
```

