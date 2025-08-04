package employee

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_employee.sql
var saveEmployeeQuery string

//go:embed sql/find_employee_by_id.sql
var findEmployeeByID string

//go:embed sql/find_employees.sql
var findEmployees string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveEmployeeStmt, err := db.Prepare(saveEmployeeQuery)
	if err != nil {
		return nil, err
	}

	findEmployeeByIDStmt, err := db.Prepare(findEmployeeByID)
	if err != nil {
		return nil, err
	}

	findEmployeesStmt, err := db.Prepare(findEmployees)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:               db,
		saveEmployee:     saveEmployeeStmt,
		findEmployeeByID: findEmployeeByIDStmt,
		findEmployees:    findEmployeesStmt,
	}, nil
}

type Repository struct {
	db               *sql.DB
	saveEmployee     *sql.Stmt
	findEmployeeByID *sql.Stmt
	findEmployees    *sql.Stmt
}
