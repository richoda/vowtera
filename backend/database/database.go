package database

import (
	"fmt"
	"log"

	"admin-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sqlx.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	if err := migrate(db); err != nil {
		log.Fatalf("Gagal migrate: %v", err)
	}

	DB = db
	log.Println("Database terhubung dan migration selesai")
}

func migrate(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id         SERIAL      PRIMARY KEY,
			name       TEXT        NOT NULL,
			email      TEXT        NOT NULL UNIQUE,
			password   TEXT        NOT NULL,
			role       TEXT        NOT NULL DEFAULT 'admin',
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		);
		CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);
	`)
	return err
}
