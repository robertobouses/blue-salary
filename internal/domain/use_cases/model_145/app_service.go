package model_145

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type Model145Repository interface {
	SaveModel145(c context.Context, model145 domain.Model145) error
}

func NewApp(Model145Repository Model145Repository) AppService {
	return AppService{
		model145Repo: Model145Repository,
	}
}

type AppService struct {
	model145Repo Model145Repository
}
