package database

import (
	"github.com/CollCaz/UniSite/database/gen/unicontentdb/public/model"
	t "github.com/CollCaz/UniSite/database/gen/unicontentdb/public/table"
	s "github.com/go-jet/jet/v2/postgres"
)

type Department struct {
	Name        string
	FacultyName string
}

type Departments []Department

func (d *DataService) scanDepartment(stmt s.Statement) (Department, error) {
	dest := model.Department{}

	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Department{}, err
	}

	dep := Department{
		Name:        dest.Name,
		FacultyName: dest.FacultyName,
	}

	return dep, nil
}

func (d *DataService) scanDepartments(stmt s.Statement) (Departments, error) {
	dest := []model.Department{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		d.logger.Info(stmt.DebugSql())
		return Departments{}, err
	}

	deps := Departments{}
	for _, dep := range dest {
		deps = append(deps, Department{
			Name:        dep.Name,
			FacultyName: dep.FacultyName,
		})
	}

	return deps, nil
}

func (d *DataService) GetAllDepartments() (Departments, error) {
	stmt := s.
		SELECT(
			t.Department.Name,
			t.Department.FacultyName,
		).FROM(t.Faculty)

	return d.scanDepartments(stmt)
}

type GetDepartmentsInFacultyArgs struct {
	FacultyName string
}

func (d *DataService) GetDepartmentsInFaculty(args GetDepartmentsInFacultyArgs) (Departments, error) {
	stmt := s.
		SELECT(
			t.Department.Name,
			t.Department.FacultyName,
		).FROM(
		t.Department,
	).WHERE(
		t.Department.FacultyName.EQ(s.String(args.FacultyName)),
	)

	return d.scanDepartments(stmt)
}

func (d *DataService) InserDepartment(dep Department) (Department, error) {
	stmt := t.Department.
		INSERT(
			t.Department.Name,
			t.Department.FacultyName,
		).VALUES(
		dep.Name,
		dep.FacultyName,
	).RETURNING(
		t.Department.Name,
		t.Department.FacultyName,
	)

	return d.scanDepartment(stmt)
}

type UpdateDepartmentArgs struct {
	old Department
	new Department
}

func (d *DataService) UpdateDepartment(args UpdateDepartmentArgs) (Department, error) {
	stmt := t.Department.
		UPDATE(
			t.Department.Name,
			t.Department.FacultyName,
		).SET(
		args.new.Name,
		args.new.FacultyName,
	).WHERE(
		s.AND(
			t.Department.Name.EQ(s.String(args.old.Name)),
			t.Department.FacultyName.EQ(s.String(args.old.FacultyName)),
		),
	).RETURNING(
		t.Department.Name,
		t.Department.FacultyName,
	)

	return d.scanDepartment(stmt)
}
