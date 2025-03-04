package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (a *app) GetAbout(c fuego.ContextNoBody) (d.AboutSection, error) {
	return a.db.GetAboutSection()
}

func (a *app) PutAbout(c fuego.ContextWithBody[d.AboutSection]) (d.AboutSection, error) {
	body, err := c.Body()
	if err != nil {
		return d.AboutSection{}, err
	}

	return a.db.UpdateAboutSection(body)
}
