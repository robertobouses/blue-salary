package employee

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_employee.sql
var saveEmployeeQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveEmployeeStmt, err := db.Prepare(saveEmployeeQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:           db,
		saveEmployee: saveEmployeeStmt,
	}, nil
}

type Repository struct {
	db           *sql.DB
	saveEmployee *sql.Stmt
}
