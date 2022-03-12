postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

startdb:
	docker start postgres12

stopdb:
	docker stop postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root birthday

dropdb:
	docker exec -it postgres12 dropdb birthday

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/birthday?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/birthday?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, sqlc, startdb, stopdb, test, server