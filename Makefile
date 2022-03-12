postgres:
	docker run --name postgres_birthday_reminders -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

startdb:
	docker start postgres_birthday_reminders

stopdb:
	docker stop postgres_birthday_reminders

createdb:
	docker exec -it postgres_birthday_reminders createdb --username=root --owner=root birthday_reminders

dropdb:
	docker exec -it postgres_birthday_reminders dropdb birthday_reminders

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/birthday_reminders?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/birthday_reminders?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go go-bday-reminders/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc startdb stopdb test server mock