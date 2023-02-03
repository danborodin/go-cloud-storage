setup:
	docker volume create --name cloud-storage
	docker network create cloud-storage
	docker compose up mongo -d
	docker exec -it mongo /bin/sh /scripts/createMongoUser.sh
	docker compose stop mongo

build:
	docker compose build

start:
	docker compose up --build -d

stop:
	docker compose stop

prune:
	docker compose down
	docker network prune -f
	docker volume prune -f