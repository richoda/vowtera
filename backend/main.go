package main

import (
	"fmt"
	"os"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	fmt.Printf("Server admin-api berjalan... [env=%s, port=%s]\n", appEnv, appPort)
}
