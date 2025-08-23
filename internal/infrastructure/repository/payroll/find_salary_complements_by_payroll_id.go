package payroll

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindSalaryComplementsByPayrollID(ctx context.Context, payrollID uuid.UUID) ([]domain.PayrollSalaryComplement, error) {
	rows, err := r.findSalaryComplementsByPayrollID.Query(payrollID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salaryComplements []domain.PayrollSalaryComplement

	for rows.Next() {
		var salaryComplement domain.PayrollSalaryComplement

		err := rows.Scan(
			&salaryComplement.ID,
			&salaryComplement.PayrollID,
			&salaryComplement.Name,
			&salaryComplement.Type,
			&salaryComplement.Value,
		)
		if err != nil {
			return nil, err
		}

		log.Printf("salary complements returned: ID=%v, PayrollID=%v, Name=%v, Type=%v, Value=%v",
			salaryComplement.ID,
			salaryComplement.PayrollID,
			salaryComplement.Name,
			salaryComplement.Type,
			salaryComplement.Value)

		salaryComplements = append(salaryComplements, salaryComplement)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return salaryComplements, nil
}
