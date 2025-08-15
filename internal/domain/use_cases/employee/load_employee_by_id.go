package employee

import (
	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) LoadEmployeeByID(employeeID uuid.UUID) (domain.Employee, error) {
	employee, err := a.employeeRepo.FindEmployeeByID(employeeID)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil

}
