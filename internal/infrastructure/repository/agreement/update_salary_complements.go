package agreement

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) UpdateSalaryComplements(ctx context.Context, agreementID uuid.UUID, complements []domain.SalaryComplement) error {
	_, err := r.deleteSalaryComplements.ExecContext(ctx, agreementID)
	if err != nil {
		return fmt.Errorf("error deleting salary complements for agreement %s: %w", agreementID, err)
	}

	for _, complement := range complements {
		_, err := r.saveSalaryComplements.ExecContext(
			ctx,
			complement.Name,
			complement.Type,
			complement.Value,
			agreementID,
		)
		if err != nil {
			return fmt.Errorf("error inserting salary complement %s: %w", complement.ID, err)
		}
	}

	return nil
}
