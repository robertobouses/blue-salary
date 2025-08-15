package payroll

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_payroll_incident.sql
var savePayrollIncidentQuery string

//go:embed sql/save_payroll.sql
var savePayrollQuery string

//go:embed sql/save_payroll_salary_complement.sql
var savePayrollSalaryComplementQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	savePayrollIncidentStmt, err := db.Prepare(savePayrollIncidentQuery)
	if err != nil {
		return nil, err
	}
	savePayrollStmt, err := db.Prepare(savePayrollQuery)
	if err != nil {
		return nil, err
	}
	savePayrollSalaryComplementStmt, err := db.Prepare(savePayrollSalaryComplementQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                  db,
		savePayrollIncident: savePayrollIncidentStmt,
		savePayroll:         savePayrollStmt,
		savePayrollSalaryComplement: savePayrollSalaryComplementStmt,
	}, nil
}

type Repository struct {
	db                  *sql.DB
	savePayrollIncident *sql.Stmt
	savePayroll         *sql.Stmt
	savePayrollSalaryComplement *sql.Stmt
}
