package agreement

import (
	"context"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type AgreementRepository interface {
	SaveAgreement(c context.Context, agreement domain.Agreement) error
	SaveCategory(c context.Context, category domain.Category) error
	UpdateAgreement(c context.Context, agreement domain.Agreement) error
	UpdateCategories(c context.Context, agreementID uuid.UUID, categories []domain.Category) error
	UpdateComplements(c context.Context, agreementID uuid.UUID, complements []domain.SalaryComplement) error
}

func NewApp(AgreementRepository AgreementRepository) AppService {
	return AppService{
		agreementRepo: AgreementRepository,
	}
}

type AppService struct {
	agreementRepo AgreementRepository
}
