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

func (d *DataService) scanFaculty(stmt s.Statement) (Faculty, error) {
	dest := model.Faculty{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Faculty{}, err
	}

	faculty := Faculty{
		Name:         dest.Name,
		Abbreviation: dest.Abbreviation,
	}

	return faculty, nil
}

func (d *DataService) scanFaculties(stmt s.Statement) (Faculties, error) {
	dest := []model.Faculty{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
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

type Faculties []Faculty

func (d *DataService) GetAllFaculties() (Faculties, error) {
	stmt := s.
		SELECT(t.Faculty.Name, t.Faculty.Abbreviation).
		FROM(t.Faculty)

	return d.scanFaculties(stmt)
}

func (d *DataService) InsertFaculty(f Faculty) (Faculty, error) {
	stmt := t.Faculty.
		INSERT(
			t.Faculty.Name,
			t.Faculty.Abbreviation,
		).VALUES(
		f.Name,
		f.Abbreviation,
	).RETURNING(
		t.Faculty.Name,
		t.Faculty.Abbreviation,
	)

	return d.scanFaculty(stmt)
}

type UpdateFaculty struct {
	Name    string
	Faculty Faculty
}

func (d *DataService) UpdateFaculty(args UpdateFaculty) (Faculty, error) {
	stmt := t.Faculty.
		UPDATE(
			t.Faculty.Name,
			t.Faculty.Abbreviation,
		).SET(
		args.Faculty.Name,
		args.Faculty.Abbreviation,
	).WHERE(
		t.Faculty.Name.EQ(s.String(args.Name)),
	).RETURNING(
		t.Faculty.Name,
		t.Faculty.Abbreviation,
	)

	return d.scanFaculty(stmt)
}
