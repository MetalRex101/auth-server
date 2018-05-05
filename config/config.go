package config

import (
	"github.com/spf13/viper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path"
)

type Config struct {
	DB *DBConfig
	App *AppConfig
}

type DBConfig struct {
	Server       string
	DSN  	     string
	Charset      string
}

type AppConfig struct {
	Env string
	MigrationDir string
}

const configPath = ".config/github.com/MetalRex101/auth-server"
const projectPath = "src/github.com/MetalRex101/auth-server"
const configName = "config"
const migrationDir = "migrations"

func GetConfig(appEnv string) *Config {
	configDir := path.Join(os.Getenv("HOME"), configPath)
	projectDir := path.Join(os.Getenv("GOPATH"), projectPath)
	migrationDir := path.Join(projectDir, migrationDir)

	os.MkdirAll(configDir, os.ModePerm)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(projectDir)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
		setDefaults()
	}

	return &Config{
		DB: &DBConfig{
			Server:       "mysql",
			DSN: 	      viper.GetString(fmt.Sprintf("%s.datasource", appEnv)),
			Charset:      "utf8",
		},
		App: &AppConfig{
			Env: appEnv,
			MigrationDir: migrationDir,
		},
	}
}

// If no config is found, set the default(s)
func setDefaults() {
	// Mysql defaults
	viper.SetDefault("db_host", "127.0.0.1")
	viper.SetDefault("db_database", "api-server")
	viper.SetDefault("db_username", "root")
	viper.SetDefault("db_password", 123456)
	viper.SetDefault("db_port", 3306)
}