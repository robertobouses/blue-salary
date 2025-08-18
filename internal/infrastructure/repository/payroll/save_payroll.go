package payroll

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SavePayroll(ctx context.Context, payroll *domain.Payroll) error {
	err := r.savePayroll.QueryRowContext(
		ctx,
		payroll.EmployeeID,
		payroll.StartDate,
		payroll.EndDate,
		payroll.ExtraPayment,
		payroll.BaseSalary,
		payroll.PersonalComplement,
		payroll.ExtraHourPay,
		payroll.MonthlyGrossWithExtras,
		payroll.BCCC,
		payroll.BCCP,
		payroll.IrpfAmount,
		payroll.IrpfEffectiveRate,
		payroll.SSContributions,
		payroll.NetSalary,
	).Scan(&payroll.ID)

	if err != nil {
		log.Print("Error executing SavePayroll statement:", err)
		return err
	}
	log.Printf("repository: payroll saved successfully with generated ID=%s", payroll.ID)
	return nil
}
