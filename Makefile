CONTAINER_NAME ?= linebot

build:
	docker build --tag ${CONTAINER_NAME} .

up:
	docker-compose up -d

down:
	docker-compose down

