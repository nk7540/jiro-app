package main

import (
	"flag"
	"log"

	"artics-api/src/config"
	"artics-api/src/config/routes"
	"artics-api/src/registry"
)

func main() {
	// Initialization
	configFile := flag.String("config", "config.yml", "User Config file")
	flag.Parse()
	app := &config.AppConfig{ConfigFile: *configFile}
	app.Setup()
	log.Print("Setup complete")

	// Registration
	reg := registry.NewRegistry(
		&app.Uploader,
		&app.Auth,
		&app.Mail,
		&app.Database,
		&app.RPC,
		&app.Websocket,
	)

	// Running application
	routes.Router(app.Server.App, reg)
	app.Server.ServeWithGracefulShutdown()
}
