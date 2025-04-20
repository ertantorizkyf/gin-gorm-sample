SERVICE_NAME := gin-gorm-sample
DOCKER_COMPOSE := docker-compose

up:
	$(DOCKER_COMPOSE) up --build -d

start:
	$(DOCKER_COMPOSE) up -d

stop:
	$(DOCKER_COMPOSE) down

restart:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) up --build -d

build:
	$(DOCKER_COMPOSE) build

logs:
	$(DOCKER_COMPOSE) logs -f app
