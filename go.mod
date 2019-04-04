module github.com/go-chassis/go-chassis-examples

require (
	github.com/go-chassis/auth v1.1.0
	github.com/go-chassis/go-archaius v0.14.0
	github.com/go-chassis/go-cc-client v0.5.0 // indirect
	github.com/go-chassis/go-chassis v1.4.0
	github.com/go-chassis/go-chassis-plugins v0.0.0-20190315050150-830f7788619a
	github.com/go-chassis/go-chassis-protocol v0.0.0-20190404065549-3aa5744623ad
	github.com/go-chassis/huawei-apm v0.0.0-20190315092647-1634e5f5cace
	github.com/go-mesh/mesher v1.6.1
	github.com/go-mesh/openlogging v0.0.0-20181205082104-3d418c478b2d
	github.com/gogo/googleapis v1.1.0 // indirect
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/huaweicse/auth v1.1.1-0.20190215074843-46b97a7adc3f
	github.com/huaweicse/cse-collector v0.0.0-20190218064311-6b8009138adb // indirect
	github.com/lyft/protoc-gen-validate v0.0.11 // indirect
	github.com/urfave/cli v1.20.1-0.20181029213200-b67dcf995b6a // indirect
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/grpc v1.16.0
)

replace (
	github.com/golang/oauth2 => ../golang.org/x/oauth2
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.3.5 => github.com/go-chassis/zipkin-go-opentracing v0.3.5-0.20190314023633-eb4b80508c56
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16 => github.com/golang/crypto v0.0.0-20181030102418-4d3f4d9ffa16
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/oauth2 v0.0.0-20180207181906-543e37812f10 => github.com/golang/oauth2 v0.0.0-20180207181906-543e37812f10
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/sys v0.0.0-20181031143558-9b800f95dbbc => github.com/golang/sys v0.0.0-20181031143558-9b800f95dbbc
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	google.golang.org/genproto v0.0.0-20181101192439-c830210a61df => github.com/google/go-genproto v0.0.0-20181101192439-c830210a61df
)
