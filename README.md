# kbds-ref-graphql

Run the following command to regenerate if there are any changes in schema.graphqls
```
go get github.com/99designs/gqlgen@v0.17.45 && \
go run github.com/99designs/gqlgen generate --verbose
```
Then update schema.resolvers.go for resolving properly,
and possibly data layer updates. 