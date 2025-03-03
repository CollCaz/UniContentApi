package database

import (
	"github.com/CollCaz/UniSite/database/gen/model"
	t "github.com/CollCaz/UniSite/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

type Faculty struct {
	Name         string `validate:"required"`
	Abbreviation string `validate:"required,alpha,len=3"`
}

type Faculties []Faculty

func (d *DataService) GetAllFaculties() (Faculties, error) {
	stmt := s.
		SELECT(t.Faculty.Name, t.Faculty.Abbreviation).
		FROM(t.Faculty)

	dest := []model.Faculty{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		return Faculties{}, err
	}

	faculites := Faculties{}
	for _, fac := range dest {
		faculites = append(faculites, Faculty{
			Name:         fac.Name,
			Abbreviation: fac.Abbreviation,
		})
	}

	return faculites, nil

}
