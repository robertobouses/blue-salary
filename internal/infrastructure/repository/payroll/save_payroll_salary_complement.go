package payroll

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SavePayrollSalaryComplement(ctx context.Context, payrollSalaryComplement domain.PayrollSalaryComplement) error {
	_, err := r.savePayrollSalaryComplement.ExecContext(
		ctx,
		payrollSalaryComplement.PayrollID,
		payrollSalaryComplement.Name,
		payrollSalaryComplement.Type,
		payrollSalaryComplement.Value,
	)
	if err != nil {
		log.Printf("repository: failed to save payroll salary complement: %v", err)
		return err
	}

	log.Println("repository: payroll salary complement saved successfully")
	return nil
}