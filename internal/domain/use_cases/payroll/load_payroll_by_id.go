package payroll

import (
	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) LoadPayrollByID(payrollID uuid.UUID) (domain.Payroll, error) {
	payroll, err := a.payrollRepo.FindPayrollByID(payrollID)
	if err != nil {
		return domain.Payroll{}, err
	}
	return payroll, nil
}
