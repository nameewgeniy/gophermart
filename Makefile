

up:
	@docker-compose -f build/docker-compose.yaml up -d

db.migration.make:
	@goose --dir internal/server/storage/pg/migrations create $(name) sql

