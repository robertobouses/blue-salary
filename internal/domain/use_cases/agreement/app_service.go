package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type AgreementRepository interface {
	SaveAgreement(c context.Context, agreement domain.Agreement) error
}

func NewApp(AgreementRepository AgreementRepository) AppService {
	return AppService{
		agreementRepo: AgreementRepository,
	}
}

type AppService struct {
	agreementRepo AgreementRepository
}
