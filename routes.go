package main

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
)

var s *fuego.Server

func (a *app) registerRoutes() {
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
}
