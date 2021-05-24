package firebase

import (
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

// App - Firebase app
type App struct {
	App *firebase.App
}

// NewApp - connects to firebase
func NewApp(ctx context.Context, config *firebase.Config, opts ...option.ClientOption) (*App, error) {
	app, err := firebase.NewApp(ctx, config, opts...)
	if err != nil {
		return nil, err
	}

	return &App{app}, nil
}
