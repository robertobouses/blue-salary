package agreement

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveAgreement(ctx context.Context, agreement domain.Agreement) error {
	_, err := r.saveAgreement.ExecContext(
		ctx,
		agreement.Name,
		agreement.NumberOfExtraPayments,
	)

	if err != nil {
		log.Print("Error executing SaveAgreement statement:", err)
		return err
	}

	return nil
}
