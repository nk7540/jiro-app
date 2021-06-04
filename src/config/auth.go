package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type AuthConfig struct {
	*auth.Client
	CGPServiceKeyJSON string `env:"GCP_SERVICE_KEY_JSON"`
}

func (c *AuthConfig) Setup() {
	opt := option.WithCredentialsJSON([]byte(c.CGPServiceKeyJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	ac, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	c.Client = ac
}
