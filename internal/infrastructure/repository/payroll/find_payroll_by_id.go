package payroll

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindPayrollByID(ctx context.Context, payrollID uuid.UUID) (domain.Payroll, error) {
	row := r.findPayrollByID.QueryRow(payrollID)

	var payroll domain.Payroll
	err := row.Scan(
		&payroll.ID,
		&payroll.EmployeeID,
		&payroll.StartDate,
		&payroll.EndDate,
		&payroll.ExtraPayment,
		&payroll.BaseSalary,
		&payroll.PersonalComplement,
		&payroll.ExtraHourPay,
		&payroll.MonthlyGrossWithExtras,
		&payroll.BCCC,
		&payroll.BCCP,
		&payroll.IrpfAmount,
		&payroll.IrpfEffectiveRate,
		&payroll.SSContributions,
		&payroll.NetSalary,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("FindPayrollByID: no payroll found with ID=%v", payrollID)
			return domain.Payroll{}, nil
		}
		log.Printf("FindPayrollByID error scanning row: %v", err)
		return domain.Payroll{}, err
	}

	log.Printf("FindPayrollByID returned payroll: %+v", payroll)
	return payroll, nil
}
