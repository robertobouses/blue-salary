package payroll

import (
	"context"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) LoadPayrollByID(ctx context.Context, payrollID uuid.UUID) (domain.Payroll, error) {
	payroll, err := a.payrollRepo.FindPayrollByID(ctx, payrollID)
	if err != nil {
		return domain.Payroll{}, err
	}
	return payroll, nil
}
