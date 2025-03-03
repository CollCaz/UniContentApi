package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

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
