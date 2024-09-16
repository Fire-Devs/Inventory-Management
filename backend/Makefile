go:
	air

MIGRATION_DIR=migration
.PHONY: migrate-up
migrate-up:
	cd $(MIGRATION_DIR) && goose postgres postgresql://root:root@localhost:5432/postgres up

migrate-down:
	cd $(MIGRATION_DIR) && goose postgres postgresql://root:root@localhost:5432/postgres down

migration-new:
	cd $(MIGRATION_DIR) && goose postgres create new_schema sql