package agreement

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

func (a AppService) CreateCategory(ctx context.Context, input agreement.CategoryRequest) error {
	log.Printf("usecase: creating category with name: %s, level: %d, base salary: %d, agreement ID: %s",
		input.Name, input.Level, input.BaseSalary, input.AgreementID)

	agreementID, err := uuid.Parse(input.AgreementID)
	if err != nil {
		log.Printf("usecase: invalid agreementID format: %v", err)
		return err
	}

	category := domain.Category{
		Name:        input.Name,
		Level:       input.Level,
		BaseSalary:  input.BaseSalary,
		AgreementID: agreementID,
	}

	if err := a.agreementRepo.SaveCategory(ctx, category); err != nil {
		log.Printf("usecase: failed to save category: %v", err)
		return err
	}

	log.Println("usecase: category created successfully")
	return nil
}
