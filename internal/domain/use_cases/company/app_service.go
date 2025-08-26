package company

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type CompanyRepository interface {
	SaveCompany(c context.Context, company domain.Company) error
}

func NewApp(CompanyRepository CompanyRepository) AppService {
	return AppService{
		companyRepo: CompanyRepository,
	}
}

type AppService struct {
	companyRepo CompanyRepository
}
