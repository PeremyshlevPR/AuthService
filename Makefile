PROTO_DIR=proto
PROTO_OUT_DIR=proto

PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto clean

proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	rm -f $(PROTO_OUT_DIR)/*.pb.go
