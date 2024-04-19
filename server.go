package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
	"github.com/wangkebin/kbds-ref-graphql/configmodel"
	"github.com/wangkebin/kbds-ref-graphql/graph"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func main() {

	log := zap.Must(zap.NewProduction())
	defer log.Sync()

	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Sugar().Errorf("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&configmodel.GlobalConfig)
	if err != nil {
		log.Sugar().Errorf("Environment can't be loaded: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = context.WithValue(ctx, "logger", log)
		// call next to execute the next middleware or resolver in the chain
		return next(ctx)
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	log.Fatal(http.ListenAndServe(":"+port, nil).Error())
}
