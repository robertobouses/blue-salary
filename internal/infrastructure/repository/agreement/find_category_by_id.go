package agreement

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindCategoryByID(categoryID uuid.UUID) (domain.Category, error) {
	log.Printf("repository: searching for category with ID %s", categoryID)

	row := r.findCategoryByID.QueryRow(categoryID)
	var category domain.Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Level,
		&category.BaseSalary,
		&category.AgreementID,
	)
	if err != nil {
		log.Printf("repository: failed to find category with ID %s: %v", categoryID, err)
		return domain.Category{}, fmt.Errorf("could not find category with ID %s: %w", categoryID, err)
	}

	log.Printf("repository: found category: %+v", category)
	return category, nil
}
