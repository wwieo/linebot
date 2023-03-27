.PHONY: up
CONTAINER_NAME ?= linebot

build:
	docker build --tag ${CONTAINER_NAME} .

up:
	@if [ ! -f ".env" ]; then \
		read -sp "Enter channelSecret: " channelSecret; \
		read -sp "Enter channelToken: " channelToken; \
		echo "channelSecret=$$channelSecret" >> .env; \
		echo "channelToken=$$channelToken" >> .env; \
	fi; \
	export channelSecret=$$(grep channelSecret .env | cut -d= -f2); \
	export channelToken=$$(grep channelToken .env | cut -d= -f2); \
	docker-compose up -d

down:
	docker-compose down
