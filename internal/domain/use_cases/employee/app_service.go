package employee

import (
	"context"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type EmployeeRepository interface {
	SaveEmployee(c context.Context, employee domain.Employee) error
	FindEmployeeByID(employeeID uuid.UUID) (domain.Employee, error)
	FindEmployees() ([]domain.Employee, error)
}

func NewApp(EmployeeRepository EmployeeRepository) AppService {
	return AppService{
		employeeRepo: EmployeeRepository,
	}
}

type AppService struct {
	employeeRepo EmployeeRepository
}
