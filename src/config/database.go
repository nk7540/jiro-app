package config

import (
	"database/sql"
	"fmt"
	"time"
)

type DatabaseConfig struct {
	*sql.DB
	Host     string `yaml:"host" env:"DB_HOST"`
	Username string `yaml:"username" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASS"`
	DBName   string `yaml:"db_name" env:"DB_NAME"`
	Port     int    `yaml:"port" env:"DB_PORT"`
}

func (d *DatabaseConfig) Setup() {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.DBName)

	db, err := sql.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(24 * time.Hour)

	d.DB = db
}
