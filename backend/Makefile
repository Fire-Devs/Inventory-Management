go:
	air

MIGRATION_DIR=migration
.PHONY: migrate-up
migrate-up:
	migrate -path $(MIGRATION_DIR) -database postgresql://root:root@localhost:5432/postgres?sslmode=disable up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database postgresql://root:root@localhost:5432/postgres?sslmode=disable down

migration-new:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq create_table