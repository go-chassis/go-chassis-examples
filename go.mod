module github.com/go-chassis/go-chassis-examples

require (
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/go-chassis/auth v1.1.2
	github.com/go-chassis/go-archaius v0.24.0
	github.com/go-chassis/go-chassis v1.7.7-0.20191202025124-badc15f68f1c
	github.com/go-chassis/go-chassis-plugins v0.0.0-20191202032917-6b1e66aa0a60
	github.com/go-chassis/go-chassis-protocol v0.0.0-20190404065549-3aa5744623ad
	github.com/go-chassis/huawei-apm v0.0.0-20190315092647-1634e5f5cace
	github.com/go-mesh/mesher v1.6.1
	github.com/go-mesh/openlogging v1.0.1
	github.com/huaweicse/auth v1.1.2
	github.com/lyft/protoc-gen-validate v0.0.11 // indirect
	github.com/urfave/cli v1.20.1-0.20181029213200-b67dcf995b6a // indirect
	google.golang.org/appengine v1.4.0 // indirect
)

replace github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5 => github.com/go-chassis/zipkin-go-opentracing v0.3.5-0.20190321072447-42cf74fc2a92

go 1.13
