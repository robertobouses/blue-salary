package agreement

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindAgreementByID(agreementID uuid.UUID) (domain.Agreement, error) {
	var agreement domain.Agreement

	agreementRow := r.findAgreementByID.QueryRow(agreementID)
	err := agreementRow.Scan(
		&agreement.ID,
		&agreement.Name,
		&agreement.NumberOfExtraPayments,
	)
	if err != nil {
		return domain.Agreement{}, err
	}

	catRows, err := r.findCategoriesByID.Query(agreementID)
	if err != nil {
		return domain.Agreement{}, err
	}
	defer catRows.Close()

	for catRows.Next() {
		var category domain.Category
		err := catRows.Scan(&category.ID, &category.Name, &category.Level, &category.BaseSalary)
		if err != nil {
			return domain.Agreement{}, err
		}
		agreement.Categories = append(agreement.Categories, category)
	}

	compRows, err := r.findComplementByID.Query(agreementID)
	if err != nil {
		return domain.Agreement{}, err
	}
	defer compRows.Close()

	for compRows.Next() {
		var complement domain.SalaryComplement
		err := compRows.Scan(&complement.ID, &complement.Name, &complement.Type, &complement.Value)
		if err != nil {
			return domain.Agreement{}, err
		}
		agreement.Complements = append(agreement.Complements, complement)
	}

	log.Printf("FindAgreementByID returned agreement: ID=%v, Name=%v, ExtraPayments=%v, Categories=%d, Complements=%d",
		agreement.ID, agreement.Name, agreement.NumberOfExtraPayments, len(agreement.Categories), len(agreement.Complements))

	return agreement, nil
}
