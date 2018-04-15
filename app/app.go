package app

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/MetalRex101/auth-server/config"
	"github.com/labstack/echo/middleware"
	"github.com/rubenv/sql-migrate"
	"github.com/MetalRex101/auth-server/app/controllers"
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/repositories"
	"github.com/MetalRex101/auth-server/app/db"
)

type App struct {
	Echo *echo.Echo
}

func (app *App) Initialize() {
	app.Echo = echo.New()

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())

	db := db.Init()
	app.migrateDB(db)

	managers := services.InitManagers(db)
	repos := repositories.InitRepositories(db)

	controllers := controllers.InitControllers(repos, managers)

	app.registerRoutes(controllers)
}

func (app *App) Run(port int) {
	app.Echo.Logger.Fatal(app.Echo.Start(fmt.Sprintf(":%d", port)))
}

func (app *App) migrateDB(db *gorm.DB) {
	if config.Instance.App.Env != "development" {
		migrations := &migrate.FileMigrationSource{
			Dir: config.Instance.DB.MigrationDir,
		}

		n, err := migrate.Exec(db.DB(), config.Instance.DB.Server, migrations, migrate.Up)

		if err != nil {
			app.Echo.Logger.Fatalf(fmt.Sprintf("Migration failed: %s", err))
		}

		fmt.Printf("Applied %d migrations!\n", n)
	}
}

func (app *App) registerRoutes(controllers *controllers.Controllers) {
	app.Echo.GET("authorize", func (c echo.Context) error {
		return controllers.Oauth.AuthorizeClient(c)
	})
}