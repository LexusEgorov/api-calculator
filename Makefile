IMAGE_NAME=api-calculator

fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test ./...
build:
	docker build -t $(IMAGE_NAME) .
run:
	docker run $(IMAGE_NAME) --config=configs/prod.yaml