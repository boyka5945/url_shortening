# generate go file from proto file
# Usage: ./script/generate_proto.sh

# generate go file from proto file
protoc --go_out=. --go-grpc_out=. ./proto/*.proto