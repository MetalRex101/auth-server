package config

import (
	"github.com/spf13/viper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"path/filepath"
	"os"
	"github.com/davecgh/go-spew/spew"
)

var Instance *Config

type Config struct {
	DB *DBConfig
	App *AppConfig
}

type DBConfig struct {
	Server       string
	DSN  	     string
	Charset      string
	MigrationDir string
}

type AppConfig struct {
	Env string
}

func GetConfig() *Config {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	spew.Dump(dir)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")

		// If no config is found, use the default(s)
		setDefaults()
	}

	return &Config{
		DB: &DBConfig{
			Server:       "mysql",
			DSN: 	      viper.GetString(fmt.Sprintf("%s.datasource", viper.GetString("app.env"))),
			Charset:      "utf8",
			MigrationDir: viper.GetString(fmt.Sprintf("%s.dir", viper.GetString("app.env"))),
		},
		App: &AppConfig{
			Env: viper.GetString("app.env"),
		},
	}
}

func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	if os.Getenv("app_env") == "testing" {
		viper.AddConfigPath("../")
	} else {
		viper.AddConfigPath(dir)
	}

	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")

		// If no config is found, use the default(s)
		setDefaults()
	}

	Instance = &Config{
		DB: &DBConfig{
			Server:       "mysql",
			DSN: 	      viper.GetString(fmt.Sprintf("%s.datasource", viper.GetString("app.env"))),
			Charset:      "utf8",
			MigrationDir: viper.GetString(fmt.Sprintf("%s.dir", viper.GetString("app.env"))),
		},
		App: &AppConfig{
			Env: viper.GetString("app.env"),
		},
	}
}

func setDefaults() {
	viper.SetDefault("app_env", "local")

	// Mysql defaults
	viper.SetDefault("db_host", "127.0.0.1")
	viper.SetDefault("db_database", "api-server")
	viper.SetDefault("db_username", "root")
	viper.SetDefault("db_password", 123456)
	viper.SetDefault("db_port", 3306)
}