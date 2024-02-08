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

.PHONY: down
down:
	sudo make migrations-copy-to-host 

	sudo docker compose down --remove-orphans
	

.PHONY: restart
restart:
	sudo make migrations-copy-to-host
	sudo docker compose down --remove-orphans
	sudo docker compose up --build -d \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate


.PHONY: create-migrations
create-migrations: # Create an alembic migration
	sudo docker compose exec coordinator_backend alembic revision --autogenerate -m "$(m)"

.PHONY: migrations-upgrade
migrations-upgrade: # Migrate to the latest migration
	sudo docker compose exec coordinator_backend alembic upgrade head

.PHONY: migrations-downgrade
migrations-downgrade: # Migrate down by 1 migration
	sudo docker compose exec coordinator_backend alembic downgrade -1

.PHONY: migrations-copy-to-host
migrations-copy-to-host: # Copy migrations from container to host
	sudo docker compose cp coordinator_backend:/app/src/alembic/versions  ./coordinator/src/alembic