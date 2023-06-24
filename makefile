include .env

migration_up:
	migrate -path db/migrations/ -database "postgresql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE}?sslmode=disable" -verbose up

migration_down:
	migrate -path db/migrations/ -database "postgresql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE}?sslmode=disable" -verbose down

migration_fix:
	migrate -path db/migrations/ -database "postgresql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE}?sslmode=disable" force VERSION
