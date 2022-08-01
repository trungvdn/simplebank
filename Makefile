postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres14 dropdb simple_bank
createsql:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate  
test:
	go test -v -cover  ./...
psql:
	docker exec -it postgres14 psql -U root -d simple_bank
server: 
	go run main.go
.PHONY: postgres createdb dropdb createsql migrateup migratedown sqlc psql server
