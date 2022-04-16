postgres:
	sudo docker run --name postgres14 -p 5432:5432  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres:14-alpine

createdb:
	sudo docker exec -it postgres14 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	sudo docker exec -it postgres14 psql -U postgres -c "drop database simple_bank"

migrateup:
	migrate -path db/migration -database "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

make:
	mockgen -package mockdb -destination db/mock/store.go github.com/ariefro/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
