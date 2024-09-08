postgres:
	docker run --name postgres12-albank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres123 -d postgres:14.4
run:
	docker start postgres12-albank
createdb:
	docker exec -it postgres12-albank createdb --username=root --owner=root al_bank
dropdb:
	docker exec -it postgres12-albank dropdb al_bank
migrateup:
	migrate -path db/migration -database "postgres://root:postgres123@localhost:5432/al_bank?sslmode=disable" -verbose up	
migratedown:
	migrate -path db/migration -database "postgres://root:postgres123@localhost:5432/al_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: postgres run createdb dropdb migrateup migratedown sqlc test server
