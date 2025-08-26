package payroll

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreatePayrollIncident(ctx context.Context, req PayrollIncidentRequest) error
	CalculatePayrollByEmployeeID(ctx context.Context, employeeID uuid.UUID, month time.Time) (domain.Payroll, error)
	CalculatePersonalComplementByEmployeeID(ctx context.Context, employeeIDstring uuid.UUID) (int, error)
	LoadIncidentByEmployeeID(employeeID uuid.UUID, month time.Time) ([]domain.PayrollIncident, error)
	CalculatePayrollsByMonth(ctx context.Context, month time.Time) ([]domain.Payroll, error)
	LoadPayrollByID(ctx context.Context, payrollID uuid.UUID) (domain.Payroll, error)
	GeneratePayrollPDFByID(ctx context.Context, payrollID uuid.UUID) (domain.GeneratePayrollPDFOutput, error)
	GeneratePayrollsPDFByMonth(ctx context.Context, month time.Time) ([]domain.GeneratePayrollPDFOutput, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
