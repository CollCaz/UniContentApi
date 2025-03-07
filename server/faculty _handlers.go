package server

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (s *Server) registerFacultiesSubrouteOn(parentRoute *fuego.Server) *fuego.Server {
	faculties := fuego.Group(parentRoute, "/faculties")
	fuego.Get(faculties, "", s.GetFaculties)
	fuego.Post(faculties, "", s.PostFacutly)
	fuego.Put(faculties, "", s.PutFaculty)

	return faculties
}

func (s *Server) GetFaculties(c fuego.ContextNoBody) (d.Faculties, error) {
	return s.db.GetAllFaculties()
}

func (s *Server) PostFacutly(c fuego.ContextWithBody[d.Faculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	return s.db.InsertFaculty(body)
}

func (s *Server) PutFaculty(c fuego.ContextWithBody[d.UpdateFaculty]) (d.Faculty, error) {
	body, err := c.Body()
	if err != nil {
		return d.Faculty{}, err
	}

	return s.db.UpdateFaculty(body)
}
