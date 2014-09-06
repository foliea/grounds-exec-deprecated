.PHONY: run binary build-go build-web test test-binary test-unit test-web images images-push images-pull

binary:
	hack/binary.sh

run:
	fig build ; fig up

build: build-go build-web

build-go:
	hack/make.sh build go

build-web:
	hack/make.sh build web

test: test-binary test-unit test-web

test-binary: build-go
	hack/make.sh test binary

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


