GOPATH:=$(shell go env GOPATH)
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin-mesh *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t shanlongpan/gin-mesh:1.1

.PHONY: docker-push
docker-push:
	docker push shanlongpan/gin-mesh:1.1

.PHONY: docker-image-tar
docker-image-tar:
	docker save shanlongpan/gin-mesh:1.1>gin-mesh.tar
