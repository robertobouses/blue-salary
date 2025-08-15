package model_145

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_model_145.sql
var saveModel145Query string

//go:embed sql/find_model_145_by_employee_id.sql
var findModel145ByEmployeeIDQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveModel145Stmt, err := db.Prepare(saveModel145Query)
	if err != nil {
		return nil, err
	}
	findModel145ByEmployeeIDStmt, err := db.Prepare(findModel145ByEmployeeIDQuery)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db:                       db,
		saveModel145:             saveModel145Stmt,
		findModel145ByEmployeeID: findModel145ByEmployeeIDStmt,
	}, nil
}

type Repository struct {
	db                       *sql.DB
	saveModel145             *sql.Stmt
	findModel145ByEmployeeID *sql.Stmt
}
