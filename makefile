GenProto:
	protoc -I ./protomsg/ ./protomsg/*.proto -I=. --go_out=plugins=grpc:protomsg