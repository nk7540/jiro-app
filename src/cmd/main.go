package main

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"ranker/src/config"
	"ranker/src/lib/grpc"
	"ranker/src/lib/redis"
	"ranker/src/registry"
)

func main() {
	// Common context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialization
	config.Logger()

	e, err := config.LoadEnvironment()
	if err != nil {
		log.Panic(err)
	}

	rdb, err := redis.NewClient(ctx, e.RedisHost, e.RedisPort, e.RedisDB)
	if err != nil {
		panic(err)
	}

	gc := grpc.NewClient(e.GrpcHost, e.GrpcPort)

	// Registration
	reg := registry.NewRegistry(rdb, gc, ws)

	// Running application
	r := config.Router(reg)
	if err := http.ListenAndServe(":"+e.Port, r); err != nil {
		log.Panic(err)
	}
}
