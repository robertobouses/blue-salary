package agreement

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) UpdateCategories(ctx context.Context, agreementID uuid.UUID, categories []domain.Category) error {
	_, err := r.deleteCategories.ExecContext(ctx, agreementID)
	if err != nil {
		return fmt.Errorf("error deleting categories for agreement %s: %w", agreementID, err)
	}

	for _, category := range categories {
		_, err := r.saveCategories.ExecContext(
			ctx,
			category.ID,
			category.Name,
			agreementID,
		)
		if err != nil {
			return fmt.Errorf("error inserting category %s: %w", category.ID, err)
		}
	}

	return nil
}
