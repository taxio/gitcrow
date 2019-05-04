SHELL := /bin/bash

NAME := gitcrow
ORG := taxio

IMAGE_NAME := $(ORG)/$(NAME)
IMAGE_VERSION := latest

image:
	docker build . -t $(IMAGE_NAME):$(IMAGE_VERSION)

prune:
	@docker image prune

