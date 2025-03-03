package main

import (
	d "github.com/CollCaz/UniSite/database"
	"github.com/go-fuego/fuego"
)

func (a *app) GetAllDepartments(c fuego.ContextNoBody) (d.Departments, error) {
	return a.db.GetAllDepartments()
}

func (a *app) GetDepsInFaculty(c fuego.ContextWithBody[d.GetDepartmentsInFacultyArgs]) (d.Departments, error) {
	body, err := c.Body()
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	return a.db.GetDepartmentsInFaculty(body)
}

func (a *app) PostDepartment(c fuego.ContextWithBody[d.Department]) (d.Department, error) {
	body, err := c.Body()
	if err != nil {
		a.logger.Error(err.Error())
		return d.Department{}, err
	}

	return a.db.InserDepartment(body)
}

func (a *app) PutDepartment(c fuego.ContextWithBody[d.UpdateDepartmentArgs]) (d.Department, error) {
	body, err := c.Body()
	if err != nil {
		a.logger.Error(err.Error())
		return d.Department{}, err
	}

	return a.db.UpdateDepartment(body)
}
