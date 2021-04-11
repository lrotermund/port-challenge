.PHONY: protos

protos:
	protoc -I server/pkg/api/ server/pkg/api/portservice.proto \
	--go-grpc_out=server/pkg/api/portservice/ \
	--go_out=server/pkg/api/portservice/ \
	--go-grpc_out=client/pkg/api/portservice/ \
	--go_out=client/pkg/api/portservice/