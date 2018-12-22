module github.com/go-chassis/go-chassis-examples

require (
	github.com/go-chassis/go-chassis v1.1.4-0.20181222093433-5096bb8a34e5
	github.com/go-chassis/go-chassis-plugins v0.0.0-20181108070724-c19e5b01b867
	github.com/go-mesh/mesher v1.5.2-0.20181112024919-f35794494256
	github.com/go-mesh/openlogging v0.0.0-20181205082104-3d418c478b2d
	github.com/gogo/googleapis v1.1.0 // indirect

	github.com/golang/protobuf v1.2.0
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/lyft/protoc-gen-validate v0.0.11 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd
	google.golang.org/grpc v1.14.0
)

replace (
	github.com/golang/oauth2 => ../golang.org/x/oauth2
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16 => github.com/golang/crypto v0.0.0-20181030102418-4d3f4d9ffa16

	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/oauth2 v0.0.0-20180207181906-543e37812f10 => github.com/golang/oauth2 v0.0.0-20180207181906-543e37812f10
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/sys v0.0.0-20181031143558-9b800f95dbbc => github.com/golang/sys v0.0.0-20181031143558-9b800f95dbbc
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	google.golang.org/genproto v0.0.0-20181101192439-c830210a61df => github.com/google/go-genproto v0.0.0-20181101192439-c830210a61df
)
