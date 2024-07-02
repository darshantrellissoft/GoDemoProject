# Define variables
DB_URL=postgres://trellis:Trellis123@localhost:5432/my_transaction
DB_NAME=my_transaction
DB_USER=trellis
DB_PASSWORD=Trellis123
DB_HOST=localhost
DB_PORT=5432
MIGRATE_CMD=migrate -path ./db/migration -database $(DB_URL)
PASETO_SECRET_KEY=0123456789abcdef0123456789abcdef

# Targets
.PHONY: help create_migration up down force create_db clear_db

help:
	@echo "Makefile for database migrations using migrate tool"
	@echo ""
	@echo "Usage:"
	@echo "  make create_migration name=<name>  Create a new migration file"
	@echo "  make migrate_up                    Run all up migrations"
	@echo "  make migrate_down                  Run all down migrations"
	@echo "  make force version=<version>       Force set database version"
	@echo "  make create_db                     Create the database if it does not exist"
	@echo "  make clear_db                      Clear the database"

create_migration:
ifndef name
	$(error name is required. Usage: make create_migration name=<name>)
endif
	@migrate create -ext sql -dir ./db/migration -seq $(name)

migrate_up:
	@$(MIGRATE_CMD) up

migrate_down:
	@$(MIGRATE_CMD) down

force:
ifndef version
	$(error version is required. Usage: make force version=<version>)
endif
	@$(MIGRATE_CMD) force $(version)

create_db:
	@sudo -u postgres psql -c "CREATE DATABASE $(DB_NAME)"
	@sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER)"

clear_db:
	@sudo -u postgres psql -c "DROP DATABASE IF EXISTS $(DB_NAME)"

run:
	@go run main.go
