package main

import (
	"context"
	"log"
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"artics-api/src/config"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/grpc"
	"artics-api/src/lib/mysql"
	"artics-api/src/registry"
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

	opt := option.WithCredentialsJSON([]byte(e.GCPServiceKeyJSON))
	fb, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(err)
	}

	fa, err := firebase.NewAuth(ctx, fb.App)
	if err != nil {
		log.Panic(err)
	}

	db, err := mysql.NewClient(ctx, e.MysqlUser, e.MysqlHost, e.MysqlPort, e.MysqlDB)
	if err != nil {
		log.Panic(err)
	}

	gc := grpc.NewClient(e.GrpcHost, e.GrpcPort)

	// Registration
	reg := registry.NewRegistry(fa, db, gc)

	// Running application
	r := config.Router(reg)
	if err := http.ListenAndServe(":"+e.Port, r); err != nil {
		log.Panic(err)
	}
}
