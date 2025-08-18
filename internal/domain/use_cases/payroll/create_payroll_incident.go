package payroll

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/payroll"
)

func (a AppService) CreatePayrollIncident(ctx context.Context, input payroll.PayrollIncidentRequest) error {
	log.Printf(
		"usecase: creating payroll incident for payroll_id: %s | description: %s | start_date: %v | end_date: %v",
		input.EmployeeID, input.Description, input.StartDate, input.EndDate,
	)

	employeeID, err := uuid.Parse(input.EmployeeID)
	if err != nil {
		log.Printf("usecase: invalid payroll_id format: %v", err)
		return err
	}

	incident := domain.PayrollIncident{
		EmployeeID:  employeeID,
		Description: input.Description,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
	}

	if err := a.payrollRepo.SavePayrollIncident(ctx, incident); err != nil {
		log.Printf("usecase: failed to save payroll incident: %v", err)
		return err
	}

	log.Println("usecase: payroll incident created successfully")
	return nil
}
