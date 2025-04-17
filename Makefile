DB_URL=postgres://user:pass@localhost:5432/shop?sslmode=disable
MIGRATION_FILE=./migrations/001_init.up.sql
BACKEND_BIN=./backend/server
SEED_BIN=./backend/seed

build:
	go build -o $(BACKEND_BIN) ./cmd/server
	go build -o $(SEED_BIN) ./cmd/seed

run:
	DATABASE_URL=$(DB_URL) $(BACKEND_BIN)

migrate:
	psql $(DB_URL) -f $(MIGRATION_FILE)

seed:
	DATABASE_URL=$(DB_URL) $(SEED_BIN)

reset:
	psql $(DB_URL) -c "DROP TABLE IF EXISTS cart, products, users CASCADE;" && \
	make migrate && make seed

dev: build migrate seed run

test:
	go test ./...

clean:
	rm -f $(BACKEND_BIN) $(SEED_BIN)
