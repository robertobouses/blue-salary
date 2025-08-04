package employee

import (
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) LoadEmployees() ([]domain.Employee, error) {
	employees, err := a.employeeRepo.FindEmployees()
	if err != nil {
		return []domain.Employee{}, err
	}
	return employees, nil

}
