package main

import (
	"database/sql"
	"log/slog"
	"os"

	d "github.com/CollCaz/UniSite/database"
	"github.com/charmbracelet/log"
	"github.com/go-fuego/fuego"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	godotenv.Load(".env", ".envrc")
	db := openDb()

	handler := log.New(os.Stderr)
	logger := slog.New(handler)

	s = fuego.NewServer(
		fuego.WithLogHandler(handler),
	)
	a := &app{
		logger: logger,
		db: d.NewDataService(d.NewDataServiceArgs{
			Db:     db,
			Logger: logger,
		}),
	}

	a.registerRoutes()

	s.Run()
}

func openDb() *sql.DB {
	dbString := os.Getenv("GOOSE_DBSTRING")
	db, err := sql.Open("sqlite3", dbString)
	if err != nil {
		panic("Could not open db")
	}

	err = db.Ping()
	if err != nil {
		panic("Could not ping db")
	}

	return db
}

type app struct {
	logger *slog.Logger
	db     *d.DataService
}

type MainAboutSection struct {
	Title   string
	Content string
}

func (a *app) helloWorld(c fuego.ContextNoBody) (string, error) {
	return "Hello, World!", nil
}
