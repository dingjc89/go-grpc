.PHONY: protos

protos :
	protoc -I service/ service/service.proto --go_out=service/ --go_opt=paths=source_relative --go-grpc_out=service/ --go-grpc_opt=paths=source_relative