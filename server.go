package main

import (
	"github.com/MetalRex101/auth-server/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(1323)
}

