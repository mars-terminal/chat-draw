.PHONY: run, migrate_down_postgres, migrate_up_postgres

export POSTGRESQL_URL="postgres://rustam:1234@localhost:5432/rustam?sslmode=disable"

migrate_up_postgres:
	golang-migrate -database ${POSTGRESQL_URL} -path scripts up
migrate_down_postgres:
	golang-migrate -database ${POSTGRESQL_URL} -path scripts down

run:
	rm -rf ./main && go build cmd/app/main.go && ./main

generate:
	swag init --dir ./cmd/app --parseDependency
