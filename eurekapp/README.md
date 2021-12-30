
on linux 

## run eureka server

clone https://github.com/spring-cloud-samples/eureka and exec
``` bash
./gradlew build
java -jar build/libs/eureka.jar
```
eureka server listen on http://127.0.0.1:8761

## run go chassis app with eureka plugin
clone this sample and exec
```bash
go mod tidy
go build
./eurekapp
```
then visit http://127.0.0.1:8761 you can see EUREKAPP in "Instances currently registered with Eureka"
