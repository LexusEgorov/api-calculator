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
	docker run --env DEPLOYMENT=local $(IMAGE_NAME)