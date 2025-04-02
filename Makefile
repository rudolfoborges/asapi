dev:
	@go run cmd/main.go

migrate:
	@docker compose up migration
