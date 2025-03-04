package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (a *app) GetHeroImages(c fuego.ContextWithBody[d.GetHeroImagesArgs]) (d.HeroImages, error) {
	amount := c.QueryParamInt("amount")
	return a.db.GetHeroImages(d.GetHeroImagesArgs{Amount: amount})
}

func (a *app) PostHeroImage(c fuego.ContextWithBody[d.HeroImage]) (d.HeroImage, error) {
	body, err := c.Body()
	if err != nil {
		a.logger.Error(err.Error())
		return d.HeroImage{}, err
	}

	return a.db.InsertHeroImage(body)

}
