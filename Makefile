postgres:
	docker run -d --name postgres_go_bank -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:15-alpine
drop-postgres:
	docker stop postgres_go_bank
	docker rm postgres_go_bank

createdb:
	docker exec -it postgres_go_bank createdb --username=postgres go_bank
dropdb:
	docker exec -it postgres_go_bank dropdb --username=postgres go_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown pull-sqlc sqlc