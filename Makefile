.PHONY: build
build:
	go build -o tmp/api-server cmd/main.go

.PHONY: protoc_docker
protoc_docker:
	docker build --platform linux/x86_64 -t protoc_image ./docker/gen-proto
	docker run --rm -it -v $(PWD):/app/ protoc_image:latest /bin/bash -c "make protoc"

.PHONY: protoc
protoc:
	protoc -I=/app/proto --go_out=/app/gen/api --go-grpc_out=/app/gen/api /app/proto/*.proto
