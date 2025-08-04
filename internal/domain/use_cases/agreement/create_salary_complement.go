package agreement

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

func (a AppService) CreateSalaryComplement(ctx context.Context, input agreement.SalaryComplementRequest) error {
	log.Printf(
		"usecase: creating salary complement with name: %s, type: %s, value: %d, agreement ID: %s",
		input.Name, input.Type, input.Value, input.AgreementID,
	)

	salaryComplement := domain.SalaryComplement{
		Name:        input.Name,
		Type:        input.Type,
		Value:       input.Value,
		AgreementID: input.AgreementID,
	}

	if err := a.agreementRepo.SaveSalaryComplement(ctx, salaryComplement); err != nil {
		log.Printf("usecase: failed to save salary complement: %v", err)
		return err
	}

	log.Println("usecase: salary complement saved successfully")
	return nil
}
