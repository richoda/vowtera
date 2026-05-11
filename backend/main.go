package main

import (
	"fmt"
	"log"
	"net/http"

	"admin-api/config"
	"admin-api/database"
	"admin-api/handlers"
	"admin-api/routes"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg)

	handlers.SetJWTSecret(cfg.JWTSecret)

	r := routes.Setup(cfg.JWTSecret)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server %s berjalan di %s [env=%s]", cfg.AppName, addr, cfg.AppEnv)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
