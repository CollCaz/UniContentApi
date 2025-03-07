package server

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (s *Server) registerAboutSubrouteOn(parentRoute *fuego.Server) *fuego.Server {
	aboutRoute := fuego.Group(parentRoute, "/about")
	fuego.Get(aboutRoute, "", s.GetAbout)
	fuego.Put(aboutRoute, "", s.PutAbout)

	return aboutRoute
}

func (s *Server) GetAbout(c fuego.ContextNoBody) (d.AboutSection, error) {
	return s.db.GetAboutSection()
}

func (s *Server) PutAbout(c fuego.ContextWithBody[d.AboutSection]) (d.AboutSection, error) {
	body, err := c.Body()
	if err != nil {
		return d.AboutSection{}, err
	}

	return s.db.UpdateAboutSection(body)
}
