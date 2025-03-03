package controller

import (
	"github.com/go-fuego/fuego"
)

type Hero_imagesResources struct {
	// TODO add resources
	Hero_imagesService Hero_imagesService
}

type Hero_images struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Hero_imagesCreate struct {
	Name string `json:"name"`
}

type Hero_imagesUpdate struct {
	Name string `json:"name"`
}

func (rs Hero_imagesResources) Routes(s *fuego.Server) {
	hero_imagesGroup := fuego.Group(s, "/hero_images")

	fuego.Get(hero_imagesGroup, "/", rs.getAllHero_images)
	fuego.Post(hero_imagesGroup, "/", rs.postHero_images)

	fuego.Get(hero_imagesGroup, "/{id}", rs.getHero_images)
	fuego.Put(hero_imagesGroup, "/{id}", rs.putHero_images)
	fuego.Delete(hero_imagesGroup, "/{id}", rs.deleteHero_images)
}

func (rs Hero_imagesResources) getAllHero_images(c fuego.ContextNoBody) ([]Hero_images, error) {
	return rs.Hero_imagesService.GetAllHero_images()
}

func (rs Hero_imagesResources) postHero_images(c *fuego.ContextWithBody[Hero_imagesCreate]) (Hero_images, error) {
	body, err := c.Body()
	if err != nil {
		return Hero_images{}, err
	}

	new, err := rs.Hero_imagesService.CreateHero_images(body)
	if err != nil {
		return Hero_images{}, err
	}

	return new, nil
}

func (rs Hero_imagesResources) getHero_images(c fuego.ContextNoBody) (Hero_images, error) {
	id := c.PathParam("id")

	return rs.Hero_imagesService.GetHero_images(id)
}

func (rs Hero_imagesResources) putHero_images(c *fuego.ContextWithBody[Hero_imagesUpdate]) (Hero_images, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Hero_images{}, err
	}

	new, err := rs.Hero_imagesService.UpdateHero_images(id, body)
	if err != nil {
		return Hero_images{}, err
	}

	return new, nil
}

func (rs Hero_imagesResources) deleteHero_images(c *fuego.ContextNoBody) (any, error) {
	return rs.Hero_imagesService.DeleteHero_images(c.PathParam("id"))
}

type Hero_imagesService interface {
	GetHero_images(id string) (Hero_images, error)
	CreateHero_images(Hero_imagesCreate) (Hero_images, error)
	GetAllHero_images() ([]Hero_images, error)
	UpdateHero_images(id string, input Hero_imagesUpdate) (Hero_images, error)
	DeleteHero_images(id string) (any, error)
}
