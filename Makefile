postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=varel -e POSTGRES_PASSWORD=newpassword237 -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=varel --owner=varel super_bank

dropdb:
	sudo docker exec -it postgres12 dropdb --username=varel super_bank

migrateup:
	migrate -path db/migrations -database "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable" -verbose down

migratefix:
	migrate -path db/migrations -database "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable" force 1

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc migratefi test
