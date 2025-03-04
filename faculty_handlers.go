package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (a *app) GetFaculties(c fuego.ContextNoBody) (d.Faculties, error) {
	return a.db.GetAllFaculties()
}

func (a *app) PostFacutly(c fuego.ContextWithBody[d.Faculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	return a.db.InsertFaculty(body)
}

func (a *app) PutFaculty(c fuego.ContextWithBody[d.UpdateFaculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	return a.db.UpdateFaculty(body)
}
