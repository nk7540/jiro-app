package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"artics-api/src/config"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/grpc"
	"artics-api/src/lib/i18n"
	"artics-api/src/lib/mysql"
	"artics-api/src/registry"
)

func main() {
	// Common context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	db, err := mysql.NewClient(ctx, e.MysqlUser, e.MysqlPassword, e.MysqlHost, e.MysqlPort, e.MysqlDB)
	if err != nil {
		log.Panic(err)
	}

	gc := grpc.NewClient(e.GrpcHost, e.GrpcPort)

	i18n.Init()

	// Registration
	reg := registry.NewRegistry(fa, db, gc)

	// Running application
	r := config.Router(reg)
	s := config.NewServer(e.Port, r)

	errCh := make(chan error, 1)
	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-ctx.Done():
		stop()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Stop(ctx); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		os.Exit(0)
	case err := <-errCh:
		log.Fatal(err)
		os.Exit(1)
	}
}
