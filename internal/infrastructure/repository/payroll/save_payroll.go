package payroll

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SavePayroll(ctx context.Context, payroll domain.Payroll) error {
	_, err := r.savePayroll.ExecContext(
		ctx,
		payroll.EmployeeID,
		payroll.StartDate,
		payroll.EndDate,
		payroll.ExtraPayment,
		payroll.BaseSalary,
		payroll.SalaryComplements,
		payroll.PersonalComplement,
		payroll.ExtraHourPay,
		payroll.MonthlyGrossWithExtras,
		payroll.BCCC,
		payroll.BCCP,
		payroll.IrpfAmount,
		payroll.IrpfEffectiveRate,
		payroll.SSContributions,
		payroll.NetSalary)

	if err != nil {
		log.Print("Error executing SavePayroll statement:", err)
		return err
	}

	return nil
}
