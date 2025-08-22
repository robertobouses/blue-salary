package payroll

import (
	"context"
	"log"
	"time"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) CalculatePayrollsByMonth(ctx context.Context, month time.Time) ([]domain.Payroll, error) {
	var payrolls []domain.Payroll
	employees, err := a.employeeRepo.FindEmployees()
	if err != nil {
		log.Println("Error finding employees", err)
		return []domain.Payroll{}, err
	}
	for _, employee := range employees {
		payroll, err := a.CalculatePayrollByEmployeeID(ctx, employee.ID, month)
		if err != nil {
			log.Println("Error Calculate Payroll By Employee ID in CalculatePayrollsByMonth")
			return []domain.Payroll{}, err
		}
		payrolls = append(payrolls, payroll)
	}
	return payrolls, nil
}
