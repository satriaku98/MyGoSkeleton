package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigServer *ConfigServer
	*ConfigDatabase
	*ConfigApp
}

type ConfigDatabase struct {
	dbConn string
}

type ConfigServer struct {
	Url  string
	Port string
}
type ConfigApp struct {
	OTHER_API_URL string
}

func newServerConfig() *ConfigServer {
	return &ConfigServer{
		GetConfigValue("SERVERURL"),
		GetConfigValue("SERVERPORT"),
	}
}

func newConfigApp() *ConfigApp {
	return &ConfigApp{
		OTHER_API_URL: GetConfigValue("OTHER_API_URL"),
	}
}

func (c *ConfigDatabase) PostgreConn() string {
	return c.dbConn
}

func ReadConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the Config file
	if err != nil {             // Handle errors reading the Config file
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
}

func GetConfigValue(configName string) string {
	ReadConfigFile("Config")
	return viper.GetString(configName)
}

func newPostgreConn() string {
	dbName := GetConfigValue("DBNAME")
	dbHost := GetConfigValue("DBHOST")
	dbUsername := GetConfigValue("DBUSERNAME")
	dbPassword := GetConfigValue("DBPASSWORD")
	dbPort := GetConfigValue("DBPORT")
	urlDb := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	// fmt.Println(urlDb)
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	return urlDb
}

func NewConfig() *Config {
	return &Config{
		ConfigServer: newServerConfig(),
		ConfigDatabase: &ConfigDatabase{
			newPostgreConn(),
		},
		ConfigApp: newConfigApp(),
	}
}
