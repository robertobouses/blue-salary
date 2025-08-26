package payroll

import (
	"context"
	"log"
	"time"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindPayrollsByMonth(ctx context.Context, month time.Time) ([]domain.Payroll, error) {
	rows, err := r.findPayrollsByMonth.QueryContext(ctx, month)
	if err != nil {
		log.Printf("FindPayrollsByMonth query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var payrolls []domain.Payroll
	for rows.Next() {
		var payroll domain.Payroll
		err := rows.Scan(
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
			log.Printf("FindPayrollsByMonth scan error: %v", err)
			return nil, err
		}
		payrolls = append(payrolls, payroll)
	}

	if err = rows.Err(); err != nil {
		log.Printf("FindPayrollsByMonth rows error: %v", err)
		return nil, err
	}

	log.Printf("FindPayrollsByMonth returned %d payrolls for month %v", len(payrolls), month)
	return payrolls, nil
}
