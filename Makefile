build:
		protoc --go_out=plugins=grpc:. proto/hello.proto
		go build -o bin/hello-client ./cmd/hello-client
		go build -o Dockerfile/hello-server ./cmd/hello-server