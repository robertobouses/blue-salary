package payroll

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreatePayrollIncident(ctx context.Context, req PayrollIncidentRequest) error
	CalculatePayrollByEmployeeID(ctx context.Context, employeeID string) (domain.Payroll, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
