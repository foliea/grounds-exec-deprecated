.PHONY: binary build-server build-web test test-unit test-web images images-push images-pull

binary:
	hack/binary.sh

build-server:
	hack/build.sh server

build-web:
	hack/build.sh web

test:
	hack/test.sh all

test-unit: build-server
	hack/test.sh unit

test-web: build-web
	hack/test.sh web

images:
	hack/images.sh build

images-push: images
	hack/images.sh push

images-pull:
	hack/images.sh pull


