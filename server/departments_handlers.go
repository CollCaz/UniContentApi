package server

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (s *Server) registerDepartmentsSubrouteOn(parentRoute *fuego.Server) *fuego.Server {
	depsRoute := fuego.Group(parentRoute, "/department")
	fuego.Get(depsRoute, "", s.GetAllDepartments)
	fuego.Post(depsRoute, "", s.PostDepartment)
	fuego.Put(depsRoute, "", s.PutDepartment)

	return depsRoute
}

func (s *Server) GetAllDepartments(c fuego.ContextNoBody) (d.Departments, error) {
	return s.db.GetAllDepartments()
}

func (s *Server) PostDepartment(c fuego.ContextWithBody[d.Department]) (d.Department, error) {
	body, err := c.Body()
	if err != nil {
		s.logger.Error(err.Error())
		return d.Department{}, err
	}

	return s.db.InserDepartment(body)
}

func (s *Server) PutDepartment(c fuego.ContextWithBody[d.UpdateDepartmentArgs]) (d.Department, error) {
	body, err := c.Body()
	if err != nil {
		s.logger.Error(err.Error())
		return d.Department{}, err
	}

	return s.db.UpdateDepartment(body)
}
