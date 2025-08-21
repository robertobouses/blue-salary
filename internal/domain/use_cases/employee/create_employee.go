package employee

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/employee"
)

func (a AppService) CreateEmployee(ctx context.Context, input employee.EmployeeRequest) error {
	log.Printf("usecase: creating employee with first name: %s, last name: %s and extra second last name: %s", input.FirstName, input.LastName, input.SecondLastName)

	categoryID, err := uuid.Parse(input.CategoryID)
	if err != nil {
		log.Printf("usecase: invalid categoryID format: %v", err)
		return err
	}

	employee := domain.Employee{
		FirstName:       input.FirstName,
		LastName:        input.SecondLastName,
		SecondLastName:  input.SecondLastName,
		GrossSalary:     input.GrossSalary,
		CategoryID:      categoryID,
		HireDate:        input.HireDate,
		TerminationDate: input.TerminationDate,
	}

	if err := a.employeeRepo.SaveEmployee(ctx, employee); err != nil {
		log.Printf("usecase: failed to save employee: %v", err)
		return err
	}

	log.Println("usecase: employee saved successfully")
	return nil
}
