IMAGE_NAME=api-calculator

apidoc:
	swag init -g cmd/calculator/main.go
fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test ./...
coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
build:
	docker build -t $(IMAGE_NAME) .
local:
	go run ./cmd/calculator --config=configs/local.yaml
dev:
	docker run $(IMAGE_NAME) --config=configs/local.yaml
run:
	docker run $(IMAGE_NAME) --config=configs/prod.yaml