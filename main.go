package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/CollCaz/UniSite/server"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	godotenv.Load(".env", ".envrc")
	db := openDb()

	handler := log.New(os.Stderr)
	logger := slog.New(handler)

	s := server.InitServer(server.NewServerArgs{
		Logger: logger,
		Db:     db,
	})

	s.Run()
}

func openDb() *sql.DB {
	dbString := os.Getenv("GOOSE_DBSTRING")
	db, err := sql.Open("pgx", dbString)
	if err != nil {
		panic("Could not open db")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("Could not ping db")
	}

	return db
}
