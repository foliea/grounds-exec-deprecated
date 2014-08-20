.PHONY: build-server build-web run run-server run-web test test-unit test-web images images-push images-pull

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)

# to allow `make REPOSITORY=custom`
REPOSITORY := $(if $(REPOSITORY),$(REPOSITORY),grounds)
# to allow `make WEB_PORT=4000`
WEB_PORT := $(if $(WEB_PORT),$(WEB_PORT),3000)
# to allow `make SERVER_PORT=4242`
SERVER_PORT := $(if $(SERVER_PORT),$(SERVER_PORT),8080)

SERVER_IMAGE := $(REGISTRY)/server$(if $(GIT_BRANCH),:$(GIT_BRANCH))
WEB_IMAGE := $(REGISTRY)/web$(if $(GIT_BRANCH),:$(GIT_BRANCH))
EXEC_IMAGES := $(shell find dockerfiles -maxdepth 1 -type d | grep dockerfiles/)

# to allow `make DOCKER_HOST=tcp://192.168.59.103:2375`
DOCKER_HOST := $(if $(DOCKER_HOST),$(DOCKER_HOST),tcp://127.0.0.1:4243)
DOCKER_IP_PORT := $(word 2, $(subst /, ,$(DOCKER_HOST)))
DOCKER_IP := $(word 1, $(subst :, ,$(DOCKER_IP_PORT)))
DOCKER_PORT := $(word 2, $(subst :, ,$(DOCKER_IP_PORT)))

binary:
	hack/binary.sh

build-server:
	docker build -t $(SERVER_IMAGE) .

build-web:
	docker build -t $(WEB_IMAGE) web

run: run-server run-web

run-redis:
	docker run -d -p 6379:6379 -v /grounds-redis:/data --name redis dockerfile/redis

run-server: build-server
	docker run -d -p $(SERVER_PORT):$(SERVER_PORT) $(SERVER_IMAGE) hack/run.sh '-d -e $(DOCKER_HOST) -r $(REPOSITORY)'

run-web: build-web
	docker run -d -p $(WEB_PORT):$(WEB_PORT) -e "RUN_ENDPOINT=$(DOCKER_IP):$(SERVER_PORT)/run RAILS_ENV=production" $(WEB_IMAGE)  rails s -p $(WEB_PORT)	

test: test-unit test-web

test-unit: build-server
	docker run --rm $(SERVER_IMAGE) hack/test-unit.sh

test-web: build-web
	docker run --rm $(WEB_IMAGE) RAILS_ENV=test bundle exec rspec
				
images:
	$(call each_exec_images,docker build -t, dir)
				
images-push: images
	$(call each_exec_images,docker push, dir)

images-pull:
	$(call each_exec_images,docker pull)

each_exec_images = $(foreach IMAGE_DIR,$(EXEC_IMAGES), \
									 $(1) $(REPOSITORY)/$(word 2, $(subst /, ,$(IMAGE_DIR))) $(if $(2),$(IMAGE_DIR));)


