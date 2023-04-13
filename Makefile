.PHONY: run
GO := $(shell which go)

run: 
	./run.sh

pb-file:
	@protoc delivery/grpc/file.proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=.	
