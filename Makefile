SHELL := /bin/bash

NAME := gitcrow
ORG := taxio

IMAGE_NAME := $(ORG)/$(NAME)
IMAGE_VERSION := latest

start:
	docker-compose up -d

start-db:
	docker-compose up -d db

stop-db:
	docker stop gitcrow-db

rm-db: stop-db
	docker rm gitcrow-db
	docker volume rm gitcrow_gitcrow-db

image:
	docker build . -t $(IMAGE_NAME):$(IMAGE_VERSION)

prune:
	@docker image prune

rm-gitcrow-dir:
	rm -rf .gitcrow/

build:
	go build

wire:
	wire gen ./app/di
