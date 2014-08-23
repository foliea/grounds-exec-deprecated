.PHONY: binary build-go build-web test test-unit test-web images images-push images-pull

binary:
	hack/binary.sh

build: build-go build-web

build-go:
	hack/make.sh build go

build-web:
	hack/make.sh build web

test: test-unit test-web

test-unit: build-go
	hack/make.sh test unit

test-web: build-web
	hack/make.sh test web

images:
	hack/make.sh images build

images-push: images
	hack/make.sh images push

images-pull:
	hack/make.sh images pull


