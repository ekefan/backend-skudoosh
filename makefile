postgresContainer:
	docker run --name postgres12 --network skudoosh-net -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWOR=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root skudoosh

dropdb:
	docker exec -it postgres12 dropdb skudoosh

makeSchema:
	migrate create -ext sql -dir ./db/migrations -seq init

migrateup:
	migrate -path ./internal/db/migrations -database "postgres://root:AFryBQNgglrAUjUTtpsOlRbhPkHb0hpJ@dpg-cpm8famehbks73da85dg-a.oregon-postgres.render.com/skudoosh" -verbose up

migratedown:
	migrate -path ./internal/db/migrations -database "postgres://root:AFryBQNgglrAUjUTtpsOlRbhPkHb0hpJ@dpg-cpm8famehbks73da85dg-a.oregon-postgres.render.com/skudoosh" -verbose down

sqlc:
	sqlc generate

force:
	migrate -path ./internal/db/migrations -database "postgres://root:AFryBQNgglrAUjUTtpsOlRbhPkHb0hpJ@dpg-cpm8famehbks73da85dg-a.oregon-postgres.render.com/skudoosh" force 1

.PHONY: postgresContainer createdb dropdb migrateup migratedown sqlc force  makeSchema
 