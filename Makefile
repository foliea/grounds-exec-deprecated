.PHONY: binary build-go build-web test test-unit test-web images images-push images-pull

binary:
	hack/binary.sh

build-go:
	hack/make.sh build go

build-web:
	hack/make.sh build web

test:
	hack/make.sh test all

test-unit: build-server
	hack/make.sh test unit

test-web: build-web
	hack/make.sh test web

images:
	hack/make.sh images build

images-push: images
	hack/make.sh images push

images-pull:
	hack/make.sh images pull


