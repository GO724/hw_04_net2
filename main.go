package main

import (
	"context"
	application "hw_04_net2/app"
)

func main() {
	app, err := application.New()
	if err != nil {
		app.Log.Fatalf("application init error: %v", err)
	}

	ctx := context.Background()
	app.Bot.Start(ctx)
}
