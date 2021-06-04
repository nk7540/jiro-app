package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"artics-api/src/config"
	"artics-api/src/registry"
)

func main() {
	// Common context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialization
	configFile := flag.String("config", "config.yml", "User Config file")
	flag.Parse()
	app := &config.AppConfig{ConfigFile: *configFile}
	app.Setup()

	// Registration
	reg := registry.NewRegistry(
		&app.Uploader,
		&app.Auth,
		&app.Mail,
		&app.Database,
		&app.RPC,
	)

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
