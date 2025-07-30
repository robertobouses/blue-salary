package model_145

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_model_145.sql
var saveModel145Query string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveModel145Stmt, err := db.Prepare(saveModel145Query)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:           db,
		saveModel145: saveModel145Stmt,
	}, nil
}

type Repository struct {
	db           *sql.DB
	saveModel145 *sql.Stmt
}
