include .env
export

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(DATABASE_URL)' up
.PHONY: migrate-up

migrate-down:
	migrate -path migrations -database '$(DATABASE_URL)' down