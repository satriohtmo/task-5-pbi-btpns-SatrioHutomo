package database_config

import "os"

var DB_DRIVER = "" 
var DB_HOST = ""
var DB_PORT = ""
var DB_NAME = ""
var DB_USER = ""
var DB_PASSWORD = ""

func InitDatabaseConfig()  {
	driverEnv := os.Getenv("DB_DRIVER")

	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")

	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")

	if portEnv != "" {
		DB_PORT = portEnv
	}

	nameEnv := os.Getenv("DB_NAME")

	if nameEnv != "" {
		DB_NAME = nameEnv
	}

	userEnv := os.Getenv("DB_USER")

	if userEnv != "" {
		DB_USER = userEnv
	}

	passwordEnv := os.Getenv("DB_PASSWORD")

	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}