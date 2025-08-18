package payroll

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreatePayrollIncident(ctx context.Context, req PayrollIncidentRequest) error
	CalculatePayrollByEmployeeID(ctx context.Context, employeeID string) (domain.Payroll, error)
	CalculatePersonalComplementByEmployeeID(ctx context.Context, employeeIDstring string) (int, error)
	LoadIncidentByEmployeeID(employeeID uuid.UUID, month time.Time) ([]domain.PayrollIncident, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
