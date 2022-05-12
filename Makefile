say_hello:
	echo "Hello World"

docker_pull_postgres:
	docker pull postgres:14.2-alpine

postgres_init:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

postgres_shell_login:
	docker exec -it postgres14 psql -U root

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simply_sacco

dropdb:
	docker exec -it postgres14 dropdb simply_sacco

create_migrations:
	migrate create -ext sql -dir db/migrations -seq init_schema

run_migrations_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simply_sacco?sslmode=disable" -verbose up

run_migrations_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simply_sacco?sslmode=disable" -verbose down

sqlc:
	sqlc generate

run_unit_tests:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: say_hello docker_pull_postgres postgres_init postgres_shell_login createdb dropdb create_migrations run_migrations_up run_migrations_down sqlc run_unit_tests server