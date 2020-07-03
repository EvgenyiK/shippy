module github.com/EvgenyiK/shippy/shippy-cli-consignment

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/EvgenyiK/shippy/shippy-service-consignment v0.0.0-20200703171029-22b958b27051
	
	github.com/micro/go-micro/v2 v2.9.1
)
