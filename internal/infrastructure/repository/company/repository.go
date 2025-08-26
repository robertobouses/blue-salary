package company

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_company.sql
var saveCompanyQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveCompanyStmt, err := db.Prepare(saveCompanyQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:          db,
		saveCompany: saveCompanyStmt,
	}, nil
}

type Repository struct {
	db          *sql.DB
	saveCompany *sql.Stmt
}
