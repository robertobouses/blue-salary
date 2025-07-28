package agreement

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_agreement.sql
var saveAgreementQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveAgreementStmt, err := db.Prepare(saveAgreementQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:            db,
		saveAgreement: saveAgreementStmt,
	}, nil
}

type Repository struct {
	db            *sql.DB
	saveAgreement *sql.Stmt
}
