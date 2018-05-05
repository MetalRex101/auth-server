package app

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/MetalRex101/auth-server/config"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/db"
	"github.com/MetalRex101/auth-server/app/handlers"
)

type App struct {
	Echo *echo.Echo
	DB *gorm.DB
	Config *config.Config
	Managers *services.Managers
	Handlers *handlers.Handlers
}

func NewApp(config *config.Config) *App {
	return &App{Config:config}
}

func (app *App) Initialize() {
	app.Echo = echo.New()

	app.InitializeServices()

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())

	app.registerRoutes()
}

func (app *App) InitializeServices () {
	app.DB = db.Init(app.Config)

	app.Managers = services.InitManagers(app.DB)

	app.Handlers = handlers.InitHandlers(
		app.Managers,
	)
}

func (app *App) Run(port int) {
	app.Echo.Logger.Fatal(app.Echo.Start(fmt.Sprintf(":%d", port)))
}

func (app *App) registerRoutes() {
	app.Echo.GET("authorize", app.Handlers.Oauth.AuthorizeClientHandler.Handle)
	app.Echo.GET("access_token", app.Handlers.Oauth.AccessTokenHandler.Handle)
}