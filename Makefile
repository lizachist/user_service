include .env
export

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'

migrate-up: ### migration up
	migrate -path migrations -database '$(DATABASE_URL)' up

migrate-down:
	migrate -path migrations -database '$(DATABASE_URL)' down