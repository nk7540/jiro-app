package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Database   DatabaseConfig `yaml:"database"`
	Auth       AuthConfig     `yaml:"auth"`
	Uploader   UploaderConfig `yaml:"uploader"`
	RPC        RPCConfig      `yaml:"rpc"`
	Logger     LoggerConfig   `yaml:"logger"`
	Mail       MailConfig     `yaml:"mail"`
	ConfigFile string
}

func (cfg *AppConfig) Setup() {
	err := godotenv.Load()
	if err = cleanenv.ReadConfig(cfg.ConfigFile, cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	cfg.Database.Setup()
	cfg.Auth.Setup()
	cfg.Uploader.Setup()
	cfg.RPC.Setup()
	cfg.Logger.Setup()
	// cfg.Mail.Setup()
}
