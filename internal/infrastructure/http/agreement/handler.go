package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type App interface {
	CreateAgreement(ctx context.Context, req AgreementRequest) error
	UpdateFullAgreement(ctx context.Context, agreement domain.Agreement) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
