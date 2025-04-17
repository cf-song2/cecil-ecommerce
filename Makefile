DB_URL=postgres://user:pass@localhost:5432/shop?sslmode=disable
MIGRATION_FILE=./backend/migrations/001_init.up.sql

BACKEND_BIN=./backend/server
SEED_BIN=./backend/seed

FRONTEND_DIR=./frontend
BACKEND_DIR=./backend
FRONTEND_PORT=30888
BACKEND_PORT=8080

build:
	go build -o $(BACKEND_BIN) ./backend/cmd/server
	go build -o $(SEED_BIN) ./backend/cmd/seed

run-backend:
	DATABASE_URL=$(DB_URL) $(BACKEND_BIN) > backend.log 2>&1 & echo $$! > backend.pid

run-frontend:
	nginx -c $(FRONTEND_DIR)/nginx.conf -p $(FRONTEND_DIR) > frontend.log 2>&1 & echo $$! > frontend.pid

stop:
	-@kill `cat backend.pid` 2>/dev/null || true
	-@kill `cat frontend.pid` 2>/dev/null || true
	-@rm -f backend.pid frontend.pid

logs:
	tail -f backend.log frontend.log

migrate:
	psql $(DB_URL) -f $(MIGRATION_FILE)

seed:
	DATABASE_URL=$(DB_URL) $(SEED_BIN)

reset:
	psql $(DB_URL) -c "DROP TABLE IF EXISTS cart, products, users CASCADE;" && \
	make migrate && make seed

dev: clean build migrate seed run-backend run-frontend
	@echo "✅ Frontend: http://localhost:$(FRONTEND_PORT)"
	@echo "✅ Backend:  http://localhost:$(BACKEND_PORT)"

test:
	go test ./...

clean:
	@rm -f $(BACKEND_BIN) $(SEED_BIN) backend.log frontend.log backend.pid frontend.pid
