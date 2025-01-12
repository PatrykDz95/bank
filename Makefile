postgres:
<<<<<<< HEAD
	podman run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	podman exec -it postgres createdb --username=root --owner=root bank

dropdb:
	podman exec -it postgres dropdb bank
=======
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres dropdb bank
>>>>>>> a322f0cdd4bc3b1652beb546b6ff2049c0f3a219

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

<<<<<<< HEAD
server:
	go run main.go

=======
>>>>>>> a322f0cdd4bc3b1652beb546b6ff2049c0f3a219
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test