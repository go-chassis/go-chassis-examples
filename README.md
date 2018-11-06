# go-chassis-examples

this repo holds go-chassis examples

most of examples can run with Apache ServiceComb service center
launch service center with docker compose
```sh
cd go-chassis-examples
sudo docker-compose up
```

but example of kubernetes and istio requires you run example in complicated environment 
that you must deploy yourself

# Dependencies
all of examples share same dependency management in go.mod

how to use go mod(go 1.11+, experimental but a recommended way)
```shell
cd go-chassis-examples
GO111MODULE=on go mod download
#optional
GO111MODULE=on go mod vendor
```