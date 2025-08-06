package payroll

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SavePayrollIncident(ctx context.Context, payrollIncident domain.PayrollIncident) error {
	_, err := r.savePayrollIncident.ExecContext(
		ctx,
		payrollIncident.PayrollID,
		payrollIncident.Description,
		payrollIncident.StartDate,
		payrollIncident.EndDate,
	)

	if err != nil {
		log.Printf("repository: failed to save payroll incident: %v", err)
		return err
	}

	log.Println("repository: payroll incident saved successfully")
	return nil
}
