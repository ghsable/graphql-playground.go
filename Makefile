GQLGEN=github.com/99designs/gqlgen

.SILENT:

generate:
	go run $(GQLGEN) generate
