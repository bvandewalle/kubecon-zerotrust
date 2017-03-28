PROJECT_NAME := kubecon17eu-server
BUILD_NUMBER := latest
DOCKER_REGISTRY=bvandewa
DOCKER_IMAGE_NAME?=$(PROJECT_NAME)
DOCKER_IMAGE_TAG?=$(BUILD_NUMBER)

build:
	GOOS=linux GOARCH=386 go build -a -installsuffix cgo
	make package

package:
	mv server docker/server

docker_build: build
	docker \
		build \
		-t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) docker

docker_push: docker_build
	docker \
		push \
		$(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
