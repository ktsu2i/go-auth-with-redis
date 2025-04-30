build:
	docker compose up --build
up:
	docker compose up
down:
	docker compose down
stop:
	docker compose stop
destroy:
	docker compose down -v --remove-orphans
front:
	docker compose exec -it frontend sh
back:
	docker compose exec -it backend sh
db:
	docker exec -it auth_db mysql -u root -p
