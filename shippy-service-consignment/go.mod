module github.com/EvgenyiK/shippy/shippy-service-consignment

go 1.14

replace github.com/EvgenyiK/shippy/shippy-service-consignment => ../shippy-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200703130916-ccbd1d7461cd // indirect
	github.com/spf13/viper v1.6.3 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200626011028-ee7919e894b5 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)
