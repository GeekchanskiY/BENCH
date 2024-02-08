up:
	sudo docker compose up \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

up-d:
	sudo docker compose up -d \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

up-build:
	sudo docker compose up --build -d \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

logs:
	sudo docker compose logs

down:
	sudo docker compose down


.PHONY: create-migrations
create-migrations: # Create an alembic migration
	sudo docker compose exec backend alembic revision --autogenerate -m "$(m)"

.PHONY: migrations-upgrade
migrations-upgrade: # Migrate to the latest migration
	sudo docker compose exec backend alembic upgrade head

.PHONY: migrations-downgrade
migrations-downgrade: # Migrate down by 1 migration
	sudo docker compose exec backend alembic downgrade -1
