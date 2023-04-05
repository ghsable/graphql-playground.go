# graphql-playground.go
[GraphQL](https://graphql.org/) Playground written by [Go](https://go.dev/).

## Development
```
.
├── .gitignore : system
├── dev.db : human
├── go.mod : system
├── go.sum : system
├── gqlgen.yml : system
├── graph : system
│   ├── generated.go : system
│   ├── helper.go : human
│   ├── model : system
│   │   └── models_gen.go : system
│   ├── resolver.go : human
│   ├── schema.graphqls : human
│   └── schema.resolvers.go : system/human
├── LICENSE : system
├── Makefile : human
├── migrate.go : human
├── README.md : human
├── server.go : system
└── tools.go : system
```

### Library
- [gqlgen](https://github.com/99designs/gqlgen)
- [gorm](https://github.com/go-gorm/gorm)

### Database
- [SQLite](https://sqlite.org/index.html)

