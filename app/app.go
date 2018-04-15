package app

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/MetalRex101/auth-server/config"
	"github.com/labstack/echo/middleware"
	"github.com/rubenv/sql-migrate"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/MetalRex101/auth-server/app/controllers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/MetalRex101/auth-server/app/db"
)

type App struct {
	Echo *echo.Echo
	Models *Models
}

type Models struct {
	Client models.Client
}

func (app *App) Initialize() {
	app.Echo = echo.New()

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
	app.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	app.migrateDB()
	app.registerRoutes()
}

func (app *App) Run(port int) {
	app.Echo.Logger.Fatal(app.Echo.Start(fmt.Sprintf(":%d", port)))
}

func (app *App) migrateDB() {
	if config.Instance.App.Env != "development" {
		migrations := &migrate.FileMigrationSource{
			Dir: config.Instance.DB.MigrationDir,
		}

		n, err := migrate.Exec(db.Gorm.DB(), config.Instance.DB.Server, migrations, migrate.Up)

		if err != nil {
			app.Echo.Logger.Fatalf(fmt.Sprintf("Migration failed: %s", err))
		}

		fmt.Printf("Applied %d migrations!\n", n)
	}
}

func (app *App) registerRoutes() {
	app.Echo.GET("authorize", func (c echo.Context) error {
		return controllers.Oauth.AuthorizeClient(c)
	})
}