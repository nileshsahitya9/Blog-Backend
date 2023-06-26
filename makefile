include .env

migration_up:
	migrate -path db/migrations/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -verbose up

migration_down:
	migrate -path db/migrations/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -verbose down

migration_fix:
	migrate -path db/migrations/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" force VERSION
