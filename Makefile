protobufs:
	rm protos/*.pb.go && protoc --go_out=plugins=grpc:protos/ protos/thing.proto
