setup:
	docker volume create --name cloud-storage
	docker network create cloud-storage
	docker compose up mongo -d
	sleep 5s
	docker exec -it mongo /bin/sh /scripts/createMongoUser.sh
	docker compose stop mongo

start:
	docker compose up --build -d

stop:
	docker compose stop

prune:
	docker compose down -v
	docker network rm cloud-storage -f
	docker volume rm cloud-storage -f