package model_145

import (
	"context"
)

type App interface {
	CreateModel145(ctx context.Context, req Model145Request) error
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
