package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

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
