package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreateAgreement(ctx context.Context, req AgreementRequest) error
	CreateCategory(ctx context.Context, req CategoryRequest) error
	CreateSalaryComplement(ctx context.Context, req SalaryComplementRequest) error
	UpdateFullAgreement(ctx context.Context, agreement domain.Agreement, categories []domain.Category, salaryComplements []domain.SalaryComplement) error
	LoadAgreements() ([]AgreementResponse, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
