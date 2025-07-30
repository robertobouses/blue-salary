package employee

import (
	"context"
)

type App interface {
	CreateEmployee(ctx context.Context, req EmployeeRequest) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
