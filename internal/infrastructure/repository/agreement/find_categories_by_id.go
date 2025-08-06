package agreement

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindCategoriesByID(agreementID uuid.UUID) ([]domain.Category, error) {
	rows, err := r.findCategoriesByID.Query(agreementID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Level,
			&category.BaseSalary,
			&category.AgreementID,
		)
		if err != nil {
			return nil, err
		}

		log.Printf("FindCategorys returned category: ID=%v, Name=%v, Level=%v, BaseSalary=%v, AgreementID=%v",
			category.ID, category.Name, category.Level, category.BaseSalary, category.AgreementID)

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
