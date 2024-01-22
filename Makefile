IMAGE_NAME := poneding/go-apiserver-template
IMAGE_TAG := latest

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o _output/bin/go-apiserver-template_linux_amd64 cmd/main.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o _output/bin/go-apiserver-template_linux_arm64 cmd/main.go

	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMAGE_NAME}:${IMAGE_TAG} --push . -f ./hack/docker/Dockerfile

.PHONY: run
run:
	swag init -g cmd/main.go
	go run cmd/main.go