package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) UpdateFullAgreement(ctx context.Context, agreement domain.Agreement, categories []domain.Category, salaryComplements []domain.SalaryComplement) error {

	if err := a.agreementRepo.UpdateAgreement(ctx, agreement); err != nil {
		return err
	}

	if err := a.agreementRepo.UpdateCategories(ctx, agreement.ID, categories); err != nil {
		return err
	}

	if err := a.agreementRepo.UpdateSalaryComplements(ctx, agreement.ID, salaryComplements); err != nil {
		return err
	}

	return nil
}
