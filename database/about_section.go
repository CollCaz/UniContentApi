package database

import (
	"github.com/CollCaz/UniSite/database/gen/model"
	t "github.com/CollCaz/UniSite/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

type AboutSection struct {
	Title   string
	Content string
}

func (d *DataService) UpdateAboutSection(args AboutSection) (AboutSection, error) {
	stmt := t.AboutSection.INSERT(
		t.AboutSection.Title,
		t.AboutSection.Content,
	).VALUES(
		args.Title,
		args.Content,
	).RETURNING(t.AboutSection.Title, t.AboutSection.Content)

	dest := model.AboutSection{}

	d.logger.Info(stmt.Sql())
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		return AboutSection{}, err
	}

	d.logger.Info(dest.Content)

	res := AboutSection{
		Title:   dest.Title,
		Content: dest.Content,
	}

	return res, nil
}

func (d *DataService) GetAboutSection() (AboutSection, error) {
	stmt := s.SELECT(
		t.AboutSection.Title,
		t.AboutSection.Content,
	).FROM(t.AboutSection).LIMIT(1)

	dest := model.AboutSection{}

	d.logger.Info(stmt.Sql())
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		return AboutSection{}, err
	}

	d.logger.Info(dest.Content)

	res := AboutSection{
		Title:   dest.Title,
		Content: dest.Content,
	}

	return res, nil
}
