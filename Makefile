IMAGE_NAME=api-calculator

fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test ./...
build:
	docker build -t $(IMAGE_NAME) .
local:
	go run ./cmd/calculator --config=configs/local.yaml
dev:
	docker run $(IMAGE_NAME) --config=configs/local.yaml
run:
	docker run $(IMAGE_NAME) --config=configs/prod.yaml