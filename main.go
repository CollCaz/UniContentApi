package main

import (
	"database/sql"
	"log/slog"
	"os"

	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	godotenv.Load(".env", ".envrc")
	db := openDb()
	s := fuego.NewServer()

	logger := slog.Default()
	a := &app{
		logger: logger,
		db: d.NewDataService(d.NewDataServiceArgs{
			Db:     db,
			Logger: logger,
		}),
	}

	fuego.Get(s, "/", a.helloWorld)

	api := fuego.Group(s, "/api")
	{ // API Routes
		about := fuego.Group(api, "/about")
		{
			fuego.Get(about, "", a.GetAbout)
			fuego.Put(about, "", a.PutAbout)
		}

		heroImages := fuego.Group(api, "/hero_images")
		{
			fuego.Get(heroImages,
				"",
				a.GetHeroImages,
				option.QueryInt(
					"amount",
					"Number of images to fetch",
					param.Default(10),
					param.Example("latest 10 images", 10),
				),
			)
			fuego.Post(heroImages, "", a.PostHeroImage)
		}

		faculties := fuego.Group(api, "/faculties")
		{
			fuego.Get(faculties, "", a.GetFaculties)
			fuego.Post(faculties, "", a.PostFacutly)
		}
	}

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

func (a *app) GetAbout(c fuego.ContextNoBody) (d.AboutSection, error) {
	section, err := a.db.GetAboutSection()
	if err != nil {
		return d.AboutSection{}, err
	}

	return section, nil
}

func (a *app) PutAbout(c fuego.ContextWithBody[d.AboutSection]) (d.AboutSection, error) {
	body, err := c.Body()
	if err != nil {
		return d.AboutSection{}, err
	}

	newSection, err := a.db.UpdateAboutSection(body)
	if err != nil {
		return d.AboutSection{}, err
	}

	return newSection, nil
}

func (a *app) GetHeroImages(c fuego.ContextWithBody[d.GetHeroImagesArgs]) (d.HeroImages, error) {
	amount := c.QueryParamInt("amount")
	heroImages, err := a.db.GetHeroImages(d.GetHeroImagesArgs{Amount: amount})
	if err != nil {
		return nil, err
	}

	return heroImages, nil
}

func (a *app) PostHeroImage(c fuego.ContextWithBody[d.HeroImage]) (d.HeroImage, error) {
	body, err := c.Body()
	if err != nil {
		a.logger.Error(err.Error())
		return d.HeroImage{}, err
	}

	heroImage, err := a.db.InsertHeroImage(body)
	if err != nil {
		a.logger.Error(err.Error())
		return d.HeroImage{}, err
	}

	return heroImage, nil

}

func (a *app) GetFaculties(c fuego.ContextNoBody) (d.Faculties, error) {
	faculites, err := a.db.GetAllFaculties()
	if err != nil {
		return nil, err
	}

	return faculites, nil
}

func (a *app) PostFacutly(c fuego.ContextWithBody[d.Faculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	faculty, err := a.db.InsertFaculty(body)
	if err != nil {
		return d.Faculty{}, err
	}

	return faculty, nil
}

func (a *app) PutFaculty(c fuego.ContextWithBody[d.UpdateFaculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	faculty, err := a.db.UpdateFaculty(body)
	if err != nil {
		return d.Faculty{}, err
	}

	return faculty, nil
}

func (a *app) helloWorld(c fuego.ContextNoBody) (string, error) {
	return "Hello, World!", nil
}
