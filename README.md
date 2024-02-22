# GeekchanskiY's bench project

TODO:
create guide to create and apply config easily

## Services:
1 - Coordinator
an fastapi app which gets all the required info from 3rd-party services and 
stores it's states
2 - Ressearch 
Ruby app with functionality to be able to do some notes of different types, and analyze them
3 - Combinator
Gin app, which allows to do some plannings, create roadmap and analysis with heavy calculations
4 - Portfolio
Django (DRF) app with portfolio items 
5 - Frontend
React application which represents data from all 3 services


## Tech stack:
    BE:
    Ruby+Rails
    Python+FastAPI
    Python+DRF
    Golang+Gin
    FE:
    JavaScript+ReactJS
    .erb
    django-templates

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

TODO: add .env readers to both applications