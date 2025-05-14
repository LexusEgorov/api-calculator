IMAGE_NAME=api-calculator

build:
	docker build -t $(IMAGE_NAME) .
run:
	docker run --env DEPLOYMENT=local $(IMAGE_NAME)