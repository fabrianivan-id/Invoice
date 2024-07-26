fmt:
	go fmt ./...

build:
	go build cmd/main.go

run:
	go run cmd/main.go

migrate.build:
	go build -o migrate migration/main/main.go

migrate.up:
	go run migration/main/main.go up

migrate.rollback:
	go run migration/main/main.go rollback

docs-update:
	rm -rf swagger/v1
	swag fmt
	swag init -d ./cmd/,./  -o swagger/v1 --ot json,yaml --pd true

test:
	go test -coverprofile cover.out ./src/...
	go tool cover -html=cover.out

up:
	docker compose up esb-invoices

stop:
	docker compose stop

down:
	docker compose down

restart:
	docker compose restart

logs:
	docker compose logs -n 30 -f esb-invoices