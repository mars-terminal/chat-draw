.phony: migrate_down_postgres migrate_up_postgres

export POSTGRESQL_URL="postgres://postgres:mysecretpassword@192.168.1.110:5432/my_little_world?sslmode=disable"

migrate_up_postgres:
	migrate -database ${POSTGRESQL_URL} -path scripts/migrate/postgres up
migrate_down_postgres:
	migrate -database ${POSTGRESQL_URL} -path scripts/migrate/postgres down
