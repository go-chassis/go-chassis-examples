This demo as an example of how to go as a provider, java as consumers

1.Install [service-center](http://servicecomb.apache.org/release/)

2.run the price service

```shell
cd go-chassis-examples/java-call-go/price
export CHASSIS_HOME=$PWD
go run main.go
```
3.run the order service.
```shell
mvn clean install
mvn spring-boot:run
```
Or run in the IDE main function

go-chassis-examples/java-call-go/order/src/main/java/com/xxx/order/Application.java

4.open http://127.0.0.1:9081/order/1101 test and verify results