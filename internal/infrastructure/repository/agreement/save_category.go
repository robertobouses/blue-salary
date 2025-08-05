package agreement

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveCategory(ctx context.Context, category domain.Category) error {
	_, err := r.saveAgreement.ExecContext(
		ctx,
		category.Name,
		category.Level,
		category.BaseSalary,
		category.AgreementID,
	)

	if err != nil {
		log.Print("Error executing SaveCategory statement:", err)
		return err
	}

	return nil
}
