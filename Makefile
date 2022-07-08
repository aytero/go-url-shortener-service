DOCKER_COMPOSE_FILE = docker-compose.yaml
all: start

start:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

re:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

clean:
	rm -rf /home/lpeggy/data/db/*
	rm -rf /home/lpeggy/data/wp/*
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes
	docker volume rm $$(docker volume ls -q)

fclean:
	docker stop $$(docker ps -qa)
	docker rm $$(docker ps -qa)
	docker rmi $$(docker images -qa)
	docker volume rm $$(docker volume ls -q)
	docker network rm $$(docker network ls -q)

.PHONY: all start stop down re fclean clean