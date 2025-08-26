package company

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_company.sql
var saveCompanyQuery string

//go:embed sql/find_company_by_agreement_id.sql
var findCompanyByAgreementIDQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveCompanyStmt, err := db.Prepare(saveCompanyQuery)
	if err != nil {
		return nil, err
	}
	findCompanyByAgreementIDStmt, err := db.Prepare(findCompanyByAgreementIDQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                       db,
		saveCompany:              saveCompanyStmt,
		findCompanyByAgreementID: findCompanyByAgreementIDStmt,
	}, nil
}

type Repository struct {
	db                       *sql.DB
	saveCompany              *sql.Stmt
	findCompanyByAgreementID *sql.Stmt
}
