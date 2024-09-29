package main

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
	// _ "github.com/urodstvo/book-shop/libs/migrations/migrations"
)

//
//go:embed migrations/*.sql
var embedMigrations embed.FS

const driver = "sqlite3"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	wd = filepath.Join(wd, "..", "..")

	dbPath := filepath.Join(wd, "db.sqlite")

	db, err := sql.Open(driver, dbPath)
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(driver); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations", goose.WithAllowMissing()); err != nil {
		panic(err)
	}

}
