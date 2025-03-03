package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

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
