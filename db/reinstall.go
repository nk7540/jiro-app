package main

import (
	"artics-api/src/config"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cfg := &config.AppConfig{ConfigFile: "config.yml"}
	if err := cleanenv.ReadConfig(cfg.ConfigFile, cfg); err != nil {
		panic(err)
	}
	cfg.Database.Setup()
	db := cfg.Database.DB

	// Migration
	migrations := &migrate.FileMigrationSource{Dir: "db/migrations"}
	_, err := migrate.Exec(db, "mysql", migrations, migrate.Down)
	if err != nil {
		panic(err)
	}
	_, err = migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	// Seeding
	q, err := ioutil.ReadFile(getSourcePath() + "/seeds/seed.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
