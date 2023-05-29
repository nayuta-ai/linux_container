PROJECT_NAME=linux_container
IMAGE_NAME=${USER}_${PROJECT_NAME}
UBUNTU_IMAGE=${IMAGE_NAME}_ubuntu
ALPINE_IMAGE=${IMAGE_NAME}_alpine
CONTAINER_NAME=${USER}_${PROJECT_NAME}
UBUNTU_CONTAINER=${CONTAINER_NAME}_ubuntu
ALPINE_CONTAINER=${CONTAINER_NAME}_alpine
SHM_SIZE=2g
FORCE_RM=true

build:
	docker build \
		-f docker/Dockerfile \
		-t $(IMAGE_NAME) \
		--no-cache \
		--force-rm=$(FORCE_RM) \
		.

run:
	docker run \
		-dit \
		-v $(PWD):/usr/src \
		--name $(CONTAINER_NAME) \
		--rm \
		--cap-add ALL \
		--shm-size $(SHM_SIZE) \
		--privileged \
		$(IMAGE_NAME) \
		/sbin/init

exec:
	docker exec \
		-it \
		$(CONTAINER_NAME) /bin/bash

stop:
	docker stop $(CONTAINER_NAME)

build_ubuntu:
	docker build \
		-f docker_ubuntu_1904/Dockerfile \
		-t $(UBUNTU_IMAGE) \
		--no-cache \
		--force-rm=$(FORCE_RM) \
		.

run_ubuntu:
	docker run \
		-dit \
		-v $(PWD):/usr/src \
		--name $(UBUNTU_CONTAINER) \
		--rm \
		--shm-size $(SHM_SIZE) \
		--privileged \
		$(UBUNTU_IMAGE) \
		/bin/bash

exec_ubuntu:
	docker exec \
		-it \
		$(UBUNTU_CONTAINER) /bin/bash

stop_ubuntu:
	docker stop $(UBUNTU_CONTAINER)

build_alpine:
	docker build \
		-f docker_alpine/Dockerfile \
		-t $(ALPINE_IMAGE) \
		--no-cache \
		--force-rm=$(FORCE_RM) \
		.

run_alpine:
	docker run \
		-dit \
		-v $(PWD):/usr/src \
		--name $(ALPINE_CONTAINER) \
		--rm \
		--shm-size $(SHM_SIZE) \
		--privileged \
		$(ALPINE_IMAGE) \
		/bin/ash

exec_alpine:
	docker exec \
		-it \
		$(ALPINE_CONTAINER) /bin/ash

stop_alpine:
	docker stop $(ALPINE_CONTAINER)

restart: stop run
