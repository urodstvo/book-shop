package database

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urodstvo/book-shop/libs/logger"
)

func New(logger logger.Logger) *sql.DB {
	wd, err := os.Getwd()
	if err != nil {
		logger.Error("failed to get working directory: " + err.Error())
	}

	dbPath := wd
	for !strings.HasSuffix(dbPath, "book-shop") {
		dbPath = filepath.Join(dbPath, "..")
	}

	dbPath = filepath.Join(dbPath, "db.sqlite")

	logger.Info("opening database: " + dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Error("failed to open database: " + err.Error())
	}

	return db
}
