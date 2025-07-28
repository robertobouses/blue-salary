package agreement

import (
	"context"
)

type App interface {
	CreateAgreement(ctx context.Context, req AgreementRequest) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
