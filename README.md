# GeekchanskiY's bench project

TODO:
create guide to create and apply config easily

!!!
export CGO_ENABLED=1

## Services:
1 - Coordinator
FastAPI
stores it's states
2 - Ressearch 
Ruby on Rails
3 - Finance
Gin
4 - Support
Django + DRF
5 - Frontend
React
6 - Big Guy
gRPC


## Tech stack:
    BE:
    Ruby+Rails
    Python+FastAPI
    Python+DRF
    Go+Gin
    FE:
    JavaScript+ReactJS

Libraries:
  TODO: add

# Setup guide:

need 2 .env files for RoR and FastApi applications databases:
.env_db_fastapi
.env_db_rails 

Content:
POSTGRES_DB=database_name
POSTGRES_USER=database_user
POSTGRES_PASSWORD=database_password
POSTGRES_HOST=database_host
POSTGRES_PORT=5432
POSTGRES_HOST_AUTH_METHOD=trust

.env_rabbitmq
RABBITMQ_DEFAULT_USER=rmuser
RABBITMQ_DEFAULT_PASS=rmpassword


# Services
![Services](img/BENCH_schema.drawio.png)

First of all, it's a test project, so it's ok for me if something does
not work at all, or something is not implemented. The goal is to
to test different frameworks, methodologies, technologies, etc.

FastAPI service is a middleware between services.
it includes:
    kafka consumer and producer
    rabbitmq consumer and producer
    redis connection for request caching
    some descriptions for other services

Ruby on Rails service:
it includes:
    rabbitmq consumer and producer

Django app:
it includes:
    kafka consumer and producer

Gin App:
it includes:
    gRPC client
    kafka consumer and producer

React App:
    websockets and http connection
    with FastAPI