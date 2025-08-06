package agreement

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

func (a AppService) CreateSalaryComplement(ctx context.Context, input agreement.SalaryComplementRequest) error {
	log.Printf(
		"usecase: creating salary complement with name: %s, type: %s, value: %d, agreement ID: %s",
		input.Name, input.Type, input.Value, input.AgreementID,
	)

	agreementID, err := uuid.Parse(input.AgreementID)
	if err != nil {
		log.Printf("usecase: invalid agreementID format: %v", err)
		return err
	}

	salaryComplement := domain.SalaryComplement{
		Name:        input.Name,
		Type:        input.Type,
		Value:       input.Value,
		AgreementID: agreementID,
	}

	if err := a.agreementRepo.SaveSalaryComplement(ctx, salaryComplement); err != nil {
		log.Printf("usecase: failed to save salary complement: %v", err)
		return err
	}

	log.Println("usecase: salary complement saved successfully")
	return nil
}
