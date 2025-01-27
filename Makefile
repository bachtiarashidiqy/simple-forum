# Set default MYSQL_URL if not set
include .env
MYSQL_URL ?= "mysql://root:$(MYSQL_ROOT_PASSWORD)@tcp(localhost:3306)/$(MYSQL_DATABASE)"

check-migrate:
	@which migrate > /dev/null || (echo "migrate not installed, please install it first" && exit 1)

# Membuat migrasi baru dengan nama yang diberikan
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Please specify a migration name with 'make migrate-create name=<migration_name>'"; \
		exit 1; \
	fi
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

# Menjalankan migrasi "up" (apply migrations)
migrate-up: check-migrate
	@ migrate -database $(MYSQL_URL) -path scripts/migrations up

# Menjalankan migrasi "down" (rollback migrations)
migrate-down: check-migrate
	@ migrate -database $(MYSQL_URL) -path scripts/migrations down
