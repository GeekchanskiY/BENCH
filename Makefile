up:
	sudo docker compose up \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

up-build:
	sudo docker compose up --build -d \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

logs:
	sudo docker compose logs

down:
	sudo docker compose down
