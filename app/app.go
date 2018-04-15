package app

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/MetalRex101/auth-server/config"
	"github.com/labstack/echo/middleware"
	"github.com/rubenv/sql-migrate"
	"github.com/MetalRex101/auth-server/app/controllers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/repositories"
)

type App struct {
	Echo *echo.Echo
}

type Controllers struct{
	Oauth *controllers.OauthController
}

func (app *App) Initialize() {
	app.Echo = echo.New()

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())
	app.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	db := app.initDB()
	app.migrateDB(db)

	managers := services.InitManagers(db)
	repos := repositories.InitRepositories(db)

	controllers := app.initControllers(managers, repos)

	app.registerRoutes(controllers)
}

func (app *App) initDB() *gorm.DB {
	conf := config.Instance
	uri := conf.DB.DSN

	conn, err := gorm.Open(conf.DB.Server, uri)

	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}

	return conn
}

func (app *App) initControllers (managers *services.Managers, repos *repositories.Repositories) *Controllers {
	base := controllers.NewBaseController(repos, managers)
	oauth := controllers.NewOauthController(base)

	return &Controllers{
		Oauth:oauth,
	}
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

func (app *App) registerRoutes(controllers *Controllers) {
	app.Echo.GET("authorize", func (c echo.Context) error {
		return controllers.Oauth.AuthorizeClient(c)
	})
}