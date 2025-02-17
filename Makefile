# variables for the current state
DOCKER_DB_NAME=snippet
DB_NAME=code-snippet-share
DB_URL=postgresql://root:secret@localhost:5432/code-snippet-share?sslmode=disable


.PHONY: run css postgres createdb dropdb migrateup migratedown sqlc

run:
	PORT=8000 go run ./cmd/main.go

css:
	npx @tailwindcss/cli -i ./api/web/assets/styles.css -o ./api/web/assets/dist/output.css --watch

#Định nghĩa target cho Postgress
postgres:
	docker run --name ${DOCKER_DB_NAME} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Định nghĩa target để tạo database
createdb:
	docker exec -it ${DOCKER_DB_NAME} createdb --username=root --owner=root ${DB_NAME}

# Định nghĩa target để xóa database
dropdb:
	docker exec -it ${DOCKER_DB_NAME} dropdb ${DB_NAME}

#Định nghĩa target để thực hiện tạo table db/migration
migrateup:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

# Định nghĩa target để thực hiện xóa db/migration
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

#Định nghĩa sqlc
sqlc:
	sqlc generate

#build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/app_prod cmd/main.go
	@echo "compiled you application with all its assets to a single binary => bin/app_prod"