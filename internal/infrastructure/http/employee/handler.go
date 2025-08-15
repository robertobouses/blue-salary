package employee

import (
	"context"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreateEmployee(ctx context.Context, req EmployeeRequest) error
	LoadEmployeeByID(employeeID uuid.UUID) (domain.Employee, error)
	LoadEmployees() ([]domain.Employee, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
