# Makefile

# Directories and file definitions
PROTO_DIR := .
GO_PROTO_DEST := ../../common/ml_server_protogen
PY_PROTO_DEST := protos
PROTO_FILES := ml_message.proto ml_service.proto

# Go code generation
.PHONY: proto-gen-go
proto-gen-go:
	protoc -I$(PROTO_DIR) \
	       --go_out=$(GO_PROTO_DEST) \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=$(GO_PROTO_DEST) \
	       --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)

# Python code generation
.PHONY: proto-gen-python
proto-gen-python:
	python -m grpc_tools.protoc -I$(PROTO_DIR) \
	       --python_out=$(PY_PROTO_DEST) \
	       --grpc_python_out=$(PY_PROTO_DEST) \
	       $(PROTO_FILES)

# Generate both Go and Python code
.PHONY: proto-gen
proto-gen: proto-gen-go proto-gen-python
