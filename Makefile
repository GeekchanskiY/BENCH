up:
	docker compose up \
	&& docker compose run ressearch_backend bin/rails db:migrate

up-d:
	sudo docker compose up -d \
	&& sudo docker compose run ressearch_backend bin/rails db:migrate

up-build:
	docker compose up --build -d \
	&& docker compose run ressearch_backend bin/rails db:migrate

logs:
	docker compose logs

.PHONY: down
down:
	make migrations-copy-to-host 

	docker compose down --remove-orphans
	

.PHONY: restart
restart:
	make migrations-copy-to-host
	docker compose down --remove-orphans
	docker compose up --build -d \
	&& docker compose run ressearch_backend bin/rails db:migrate


.PHONY: create-migrations
create-migrations: # Create an alembic migration
	docker compose exec coordinator_backend alembic revision --autogenerate -m "$(m)"

.PHONY: migrations-upgrade
migrations-upgrade: # Migrate to the latest migration
	docker compose exec coordinator_backend alembic upgrade head

.PHONY: migrations-downgrade
migrations-downgrade: # Migrate down by 1 migration
	docker compose exec coordinator_backend alembic downgrade -1

.PHONY: migrations-copy-to-host
migrations-copy-to-host: # Copy migrations from container to host
	docker compose cp coordinator_backend:/app/src/alembic/versions  ./coordinator/src/alembic


rails-sandbox:
	docker compose exec -it ressearch_backend sh -c "bin/rails console --sandbox"

rails-console:
	docker compose exec -it ressearch_backend sh -c "bin/rails console"