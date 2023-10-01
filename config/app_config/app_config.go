package app_config

import "os"

var PORT = ""

func InitAppConfig() {
	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		PORT = portEnv
	}
}