.PHONY: protos

install:
	go get github.com/hashicorp/go-hclog && go get google.golang.org/grpc && go get google.golang.org/protobuf

protos:
	 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/currency.proto