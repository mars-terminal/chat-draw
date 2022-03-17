.PHONY: run, migrate_down_postgres, migrate_up_postgres

export POSTGRESQL_URL="postgres://${POSTGRESQL_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable"

migrate_up_postgres:
	golang-migrate -database ${POSTGRESQL_URL} -path scripts up
migrate_down_postgres:
	golang-migrate -database ${POSTGRESQL_URL} -path scripts down

run:
	rm -rf ./main && go build cmd/app/main.go && ./main

generate:
	swag init --dir ./cmd/app --parseDependency
