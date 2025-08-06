package payroll

import (
	"context"
)

type App interface {
	CreatePayrollIncident(ctx context.Context, req PayrollIncidentRequest) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
