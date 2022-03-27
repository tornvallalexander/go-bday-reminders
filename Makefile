postgres:
	docker run --name postgres-birthday-reminders --network birthday-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

up:
	docker compose up

down:
	docker compose down

startdb:
	docker start postgres-birthday-reminders

stopdb:
	docker stop postgres-birthday-reminders

createdb:
	docker exec -it postgres-birthday-reminders createdb --username=root --owner=root birthday_reminders

dropdb:
	docker exec -it postgres-birthday-reminders dropdb birthday_reminders

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

# migrate create -ext sql -dir db/migration -seq <migration_name>

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock up down startdb stopdb