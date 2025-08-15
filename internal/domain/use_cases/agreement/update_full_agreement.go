package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) UpdateFullAgreement(ctx context.Context, agreement domain.Agreement, categories []domain.Category, salaryComplements []domain.SalaryComplement) error {
	err := a.agreementRepo.UpdateAgreement(ctx, agreement)
	if err != nil {
		return err
	}

	err = a.agreementRepo.UpdateCategories(ctx, agreement.ID, categories)
	if err != nil {
		return err
	}

	err = a.agreementRepo.UpdateSalaryComplements(ctx, agreement.ID, salaryComplements)
	if err != nil {
		return err
	}

	return nil
}
