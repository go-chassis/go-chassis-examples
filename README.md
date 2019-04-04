# go-chassis-examples

this repo holds go-chassis examples


1.most of examples can run with Apache ServiceComb service center
launch service center with docker compose
```sh
cd go-chassis-examples
sudo docker-compose up
```

2.example of kubernetes and istio requires you run example in complicated environment 
that you must deploy yourself

3.example of huawei cloud, 
you need to register huawei cloud account to use https://www.huaweicloud.com/product/servicestage.html

# Dependencies
all of examples share same dependency management in go.mod. 

before you start you must download all dependencies with go mod. 

how to use go mod(go 1.11+, experimental but a recommended way)
```shell
cd go-chassis-examples
GO111MODULE=on go mod download
#optional
GO111MODULE=on go mod vendor
```


# Scenarios

## distributed configuration management

## canary release

## protect you system with circuit breaker to prevent cascade failure

## mutual TLS