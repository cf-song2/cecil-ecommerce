COMPOSE = docker compose
BACKEND = backend
DB = db

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up

upd:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

clean:
	$(COMPOSE) down -v --remove-orphans

reup:
	$(COMPOSE) down -v --remove-orphans
	$(COMPOSE) build
	$(COMPOSE) up -d

logs:
	$(COMPOSE) logs -f $(BACKEND)

sh:
	$(COMPOSE) exec $(BACKEND) /bin/sh

psql:
	$(COMPOSE) exec $(DB) psql -U user -d shop

# Do not use this manually
migrate:
	$(COMPOSE) exec $(DB) psql -U user -d shop -f /app/migrations/001_init.up.sql

seed:
	$(COMPOSE) exec $(BACKEND) ./seed

ps:
	$(COMPOSE) ps
