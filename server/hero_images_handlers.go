package server

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (s *Server) registerHeroImagesSubrouteOn(parentRoute *fuego.Server) *fuego.Server {
	heroRoute := fuego.Group(parentRoute, "/hero_images")
	fuego.Get(heroRoute, "", s.GetHeroImages)
	fuego.Post(heroRoute, "", s.PostHeroImage)

	return heroRoute
}

func (s *Server) GetHeroImages(c fuego.ContextWithBody[d.GetHeroImagesArgs]) (d.HeroImages, error) {
	amount := c.QueryParamInt("amount")
	return s.db.GetHeroImages(d.GetHeroImagesArgs{Amount: amount})
}

func (s *Server) PostHeroImage(c fuego.ContextWithBody[d.HeroImage]) (d.HeroImage, error) {
	body, err := c.Body()
	if err != nil {
		s.logger.Error(err.Error())
		return d.HeroImage{}, err
	}

	return s.db.InsertHeroImage(body)

}
