up-build:
	sudo docker-compose up --build -d \
	&& sudo docker-compose run ressearch_backend bin/rails db:migrate

down:
	sudo docker-compose down
