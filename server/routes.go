package server

import (
	"github.com/go-fuego/fuego"
)

func (s *Server) RegisterRoutes() {
	apiRoute := fuego.Group(s.server, "/api-v1")
	{
		fuego.Get(apiRoute, "/", s.helloWorld)
		// About Section
		s.registerAboutSubrouteOn(apiRoute)
		// Main Hero Images
		s.registerHeroImagesSubrouteOn(apiRoute)
		// Departments
		s.registerDepartmentsSubrouteOn(apiRoute)
		// Events
		s.registerEventsSubrouteOn(apiRoute)
		// Faculties
		s.registerFacultiesSubrouteOn(apiRoute)
	}
}

func (s *Server) helloWorld(c fuego.ContextNoBody) (string, error) {
	return "Hello World!", nil
}
