# graphql-playground.go
[GraphQL](https://graphql.org/) Playground written by [Go](https://go.dev/).

## development
```
.
├── dev.db : human
├── go.mod : system
├── go.sum : system
├── gqlgen.yml : system
├── graph : system
│   ├── generated.go : system
│   ├── model : system
│   │   └── models_gen.go : system
│   ├── resolver.go : human
│   ├── schema.graphqls : human
│   └── schema.resolvers.go : human
├── LICENSE : system
├── Makefile : human
├── migrate.go : human
├── README.md : human
├── server.go : system
└── tools.go : system
```

### library
- [gqlgen](https://github.com/99designs/gqlgen)
- [gorm](https://github.com/go-gorm/gorm)

### database
- [SQLite](https://sqlite.org/index.html)

