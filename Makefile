POSTGRES-PORT ?= 5432
POSTGRES-PASS ?= qwerty

install-go-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

run-postgres:
	docker run --name=todo-db -e POSTGRES_PASSWORD=$(POSTGRES-PASS) -p $(POSTGRES-PORT):5432 -d --rm postgres

# down
COMMAND ?= up
migrate:
	migrate -path ./migrations -database 'postgres://postgres:$(POSTGRES-PASS)@localhost:$(POSTGRES-PORT)/postgres?sslmode=disable' $(COMMAND)
