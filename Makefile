build:
	go build -o ./bin/url_shortening ./cmd/url_shortening.go

run:
	./bin/url_shortening

generate-proto:
	./script/generate_proto.sh