package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server     ServerConfig    `yaml:"server"`
	Database   DatabaseConfig  `yaml:"database"`
	Auth       AuthConfig      `yaml:"auth"`
	Uploader   UploaderConfig  `yaml:"uploader"`
	RPC        RPCConfig       `yaml:"rpc"`
	I18n       I18nConfig      `yaml:"i18n"`
	Logger     LoggerConfig    `yaml:"logger"`
	Mail       MailConfig      `yaml:"mail"`
	Websocket  WebsocketConfig `yaml:"websocket"`
	ConfigFile string
}

func (cfg *AppConfig) Setup() {
	err := godotenv.Load()
	if err = cleanenv.ReadConfig(cfg.ConfigFile, cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	cfg.Server.LoadPath()
	cfg.Server.Setup()
	cfg.Database.Setup()
	cfg.Auth.Setup()
	cfg.Uploader.Setup()
	cfg.RPC.Setup()
	cfg.I18n.Setup()
	cfg.Server.Use(cfg.I18n.NewMiddleware())
	cfg.Logger.Setup()
	// cfg.Mail.Setup()
}
