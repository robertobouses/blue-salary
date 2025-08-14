package payroll

import (
	"context"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type PayrollRepository interface {
	SavePayrollIncident(c context.Context, payrollIncident domain.PayrollIncident) error
	SavePayroll(ctx context.Context, payroll domain.Payroll) error
}

type EmployeeRepository interface {
	FindEmployeeByID(employeeID uuid.UUID) (domain.Employee, error)
}

type AgreementRepository interface {
	FindCategoryByID(categoryID uuid.UUID) (domain.Category, error)
	FindSalaryComplementsByID(agreementID uuid.UUID) ([]domain.SalaryComplement, error)
	FindAgreementByID(agreementID uuid.UUID) (domain.Agreement, error)
}

type Model145Repository interface {
	FindModel145ByEmployeeID(ctx context.Context, employeeID uuid.UUID) (domain.Model145, error)
}

func NewApp(
	payrollRepository PayrollRepository,
	employeeRepository EmployeeRepository,
	agreementRepository AgreementRepository,
	model145Repository Model145Repository,
) AppService {
	return AppService{
		payrollRepo:   payrollRepository,
		employeeRepo:  employeeRepository,
		agreementRepo: agreementRepository,
		model145Repo:  model145Repository,
	}

}

type AppService struct {
	payrollRepo   PayrollRepository
	employeeRepo  EmployeeRepository
	agreementRepo AgreementRepository
	model145Repo  Model145Repository
}
