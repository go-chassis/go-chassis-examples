module github.com/go-chassis/go-chassis-examples

require (
	github.com/apache/servicecomb-mesher v1.6.3 // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/go-chassis/auth v1.1.2
	github.com/go-chassis/go-archaius v1.2.1
	github.com/go-chassis/go-chassis v1.8.3
	github.com/go-chassis/go-chassis-extension/protocol/grpc v0.0.0-20200528045443-7466a321557f
	github.com/go-chassis/go-chassis-plugins v0.0.0-20191202032917-6b1e66aa0a60
	github.com/go-chassis/go-chassis-protocol v0.0.0-20200525011845-d1643389cdb8
	github.com/go-chassis/huawei-apm v0.0.0-20190315092647-1634e5f5cace // indirect
	github.com/go-mesh/openlogging v1.0.1
	github.com/golang/protobuf v1.3.3
	github.com/huaweicse/auth v1.1.2
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/grpc v1.29.1
)

replace github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5 => github.com/go-chassis/zipkin-go-opentracing v0.3.5-0.20190321072447-42cf74fc2a92

go 1.13
