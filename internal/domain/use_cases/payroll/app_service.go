package payroll

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type PayrollRepository interface {
	SavePayrollIncident(c context.Context, payrollIncident domain.PayrollIncident) error
	SavePayroll(ctx context.Context, payroll *domain.Payroll) error
	SavePayrollSalaryComplement(ctx context.Context, payrollSalaryComplement domain.PayrollSalaryComplement) error
	FindIncidentByEmployeeID(employeeID uuid.UUID, month time.Time) ([]domain.PayrollIncident, error)
	FindPayrollByID(ctx context.Context, payrollID uuid.UUID) (domain.Payroll, error)
	FindSalaryComplementsByPayrollID(ctx context.Context, payrollID uuid.UUID) ([]domain.PayrollSalaryComplement, error)
	FindPayrollsByMonth(ctx context.Context, month time.Time) ([]domain.Payroll, error)
}

type EmployeeRepository interface {
	FindEmployeeByID(employeeID uuid.UUID) (domain.Employee, error)
	FindEmployees() ([]domain.Employee, error)
}

type AgreementRepository interface {
	FindCategoryByID(categoryID uuid.UUID) (domain.Category, error)
	FindSalaryComplementsByID(agreementID uuid.UUID) ([]domain.SalaryComplement, error)
	FindAgreementByID(agreementID uuid.UUID) (domain.Agreement, error)
}

type Model145Repository interface {
	FindModel145ByEmployeeID(ctx context.Context, employeeID uuid.UUID) (domain.Model145, error)
}

type PDFService interface {
	RenderPayroll(payroll domain.Payroll, complements []domain.PayrollSalaryComplement) ([]byte, error)
}

func NewApp(
	payrollRepository PayrollRepository,
	employeeRepository EmployeeRepository,
	agreementRepository AgreementRepository,
	model145Repository Model145Repository,
	pdfService PDFService,
) AppService {
	return AppService{
		payrollRepo:   payrollRepository,
		employeeRepo:  employeeRepository,
		agreementRepo: agreementRepository,
		model145Repo:  model145Repository,
		pdfService:    pdfService,
	}

}

type AppService struct {
	payrollRepo   PayrollRepository
	employeeRepo  EmployeeRepository
	agreementRepo AgreementRepository
	model145Repo  Model145Repository
	pdfService    PDFService
}
