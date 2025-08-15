package agreement

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) UpdateAgreement(ctx context.Context, agreement domain.Agreement) error {
	findedAgreement, err := r.FindAgreementByID(agreement.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("agreement not found with ID: %s", agreement.ID)
		}
		return fmt.Errorf("error fetching current agreement: %w", err)
	}

	if agreement.Name == "" {
		agreement.Name = findedAgreement.Name
	}
	if agreement.NumberOfExtraPayments < 0 {
		agreement.NumberOfExtraPayments = findedAgreement.NumberOfExtraPayments
	}

	_, err = r.updateAgreement.ExecContext(
		ctx,
		agreement.Name,
		agreement.NumberOfExtraPayments,
		agreement.ID,
	)
	if err != nil {
		log.Printf("Error updating agreement: %v", err)
		return err
	}

	return nil
}
