#!bin/bash
export POSTGRESQL_URL=postgres://anbui:1234@localhost:5432/instagram_db?sslmode=disable

# make create-migrations name=add_a_column
create-migration:
	@migrate create -ext sql -dir db/migrations -seq $(name)


# make migrations-up num=1
migrations-up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up $(num)


# make migrations-down num=1
migrations-down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down $(num)

# run app
run-app:
	@go run cmd/app/main.go

# Swagger app
swagger-app:
	swag init -d ./ -g cmd/app/main.go \ -o ./docs/app

