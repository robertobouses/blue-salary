package agreement

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) UpdateFullAgreement(ctx context.Context, agreement domain.Agreement) error {
	err := a.agreementRepo.UpdateAgreement(ctx, agreement)
	if err != nil {
		return err
	}

	err = a.agreementRepo.UpdateCategories(ctx, agreement.ID, agreement.Categories)
	if err != nil {
		return err
	}

	err = a.agreementRepo.UpdateComplements(ctx, agreement.ID, agreement.Complements)
	if err != nil {
		return err
	}

	return nil
}
