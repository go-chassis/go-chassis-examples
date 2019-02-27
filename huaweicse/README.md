this demo shows how to use huawei cloud service engine as control plane. Check [doc](https://docs.go-chassis.com/huawei/cse.html)
How to run

1. input your ak sk in auth.yaml

2. check web console to get your service center address, input it in chassis.yaml
```yaml
cse:
  service:
    registry:
      address: http://127.0.0.1:30100 # service center address
```

2. 
```go
go build main.go
./main
```


For more detail about how to use huawei cloud service engine, see https://www.huaweicloud.com/product/cse.html