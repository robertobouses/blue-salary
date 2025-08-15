package agreement

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

func (a AppService) CreateAgreement(ctx context.Context, input agreement.AgreementRequest) error {
	log.Printf("usecase: creating agreement with name: %s and extra payments: %d", input.Name, input.NumberOfExtraPayments)

	agreement := domain.Agreement{
		Name:                  input.Name,
		NumberOfExtraPayments: input.NumberOfExtraPayments,
	}

	if err := a.agreementRepo.SaveAgreement(ctx, agreement); err != nil {
		log.Printf("usecase: failed to save agreement: %v", err)
		return err
	}

	log.Println("usecase: agreement saved successfully")
	return nil
}
