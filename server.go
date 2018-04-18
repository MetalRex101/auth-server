package main

import (
	"github.com/MetalRex101/auth-server/app"
	"github.com/MetalRex101/auth-server/config"
	"os"
)

func main() {
	config := config.GetConfig(os.Getenv("APP_ENV"))

	app := app.NewApp(config)
	app.Initialize()
	app.Run(1323)
}

