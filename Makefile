.PHONY: build
build:
	go build -o tmp/api-server cmd/main.go

.PHONY: protoc-api_docker
protoc-api_docker:
	docker build --platform linux/x86_64 -t protoc_image ./docker/gen-proto
	docker run --rm -it -v $(PWD):/app/ protoc_image:latest /bin/bash -c "make protoc-api"

.PHONY: protoc-api
protoc-api:
	protoc -I=/app/proto --go_out=/app/gen/pb/api --go-grpc_out=/app/gen/pb/api /app/proto/api-server/*.proto
