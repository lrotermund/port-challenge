.PHONY: protos

protos:
	protoc -I server/pkg/api/v1/ server/pkg/api/v1/portservice.proto \
	--go-grpc_out=server/pkg/api/v1/portservice/ \
	--go_out=server/pkg/api/v1/portservice/ \
	--go-grpc_out=client/pkg/api/v1/portservice/ \
	--go_out=client/pkg/api/v1/portservice/