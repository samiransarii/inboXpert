# samiransari@macbook:~/D/p/s/backend|main⚡*?
# ➤ protoc -I. -I../ \           (base) 
#         --go-grpc_out=. \
#         --go-grpc_opt=paths=source_relative \
#         proto/services/email_categorization_service.proto


# samiransari@macbook:~/D/p/s/backend|main⚡*?
# protoc -I. \           (base) 
#           --go_out=. \
#           --go_opt=paths=source_relative \
#           proto/messages/email_categorization.proto

# Makefile for generating protobuf files

PROTO_SRC_DIR=proto/services
PROTO_GEN_DIR=proto/messages
OUT_DIR=.

.PHONY: all grpc proto

# Run both grpc and proto commands
all: grpc proto

# Generate gRPC code
grpc:
	protoc -I. -I../ \
		--go-grpc_out=$(OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC_DIR)/email_categorization_service.proto

# Generate Go code
proto:
	protoc -I. \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		$(PROTO_GEN_DIR)/email_categorization.proto
