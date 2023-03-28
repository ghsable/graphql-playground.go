GQLGEN=github.com/99designs/gqlgen

.SILENT:

init:
	go run $(GQLGEN) init

generate:
	go run $(GQLGEN) generate

migrate:
	go run migrate.go

server:
	go run server.go
