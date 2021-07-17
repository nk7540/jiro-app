package main

import (
	"artics-api/src/config"
	"artics-api/src/internal/models"
	"context"
	"flag"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()

	configFile := flag.String("config", "config.yml", "User Config file")
	flag.Parse()
	app := &config.AppConfig{ConfigFile: *configFile}
	app.Setup()

	db := app.Database.DB

	// User
	u := &models.User{
		UID:          "LrrvRiLQ9ed2hdwHqu9k12HEWKh2",
		Status:       "provisional",
		Email:        "A@g.com",
		Nickname:     "にっくねーむ",
		Profile:      "ぷろふぃーる",
		ThumbnailURL: "https://artics-s3.s3.ap-northeast-1.amazonaws.com/1626321170037.jpg",
	}
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}

	categoryNames := []string{"comic", "anime", "music", "drama", "movie"}
	for _, name := range categoryNames {
		c := models.Category{Name: name}
		if err := c.Insert(ctx, db, boil.Infer()); err != nil {
			panic(err)
		}
	}
}
