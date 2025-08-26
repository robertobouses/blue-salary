package company

import (
	"context"
)

type App interface {
	CreateCompany(ctx context.Context, input CompanyRequest) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
