## Regular Expression

go-chassis rest protocol support path parameter and dynamic url


1 path parameter

set path parameter for rest protocol ,like
```go
[]restful.Route{
	{Method: "GET", Path: "/param/{path_parameter}", ResourceFuncName: "Func_Name"},
}
``` 
The above example shows that `path_parameter` is the name of path parameter.you can get value through it.


2 dynamic router

use regex to implement dynamic router , you can match all access like 
```go
[]restful.Route{
	{Method: "GET", Path: "/dynamic/{dynamic:*}", ResourceFuncName: "Func_Name"},
}
``` 

The above example shows can match all call which url prefix is `/dynamic` .
If you did not match all call , you can only match letter with regex, like
```go
[]restful.Route{
	{Method: "GET", Path: "/dynamic/{dynamic:[a-zA-Z]}", ResourceFuncName: "Func_Name"},
}
``` 