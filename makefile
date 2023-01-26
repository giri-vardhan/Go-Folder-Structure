migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

migrate:
	migrate create -ext sql -dir db/migration -seq ${FILE_NAME}

server:
	go run main.go