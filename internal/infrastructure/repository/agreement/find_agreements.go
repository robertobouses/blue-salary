package agreement

import (
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindAgreements() ([]domain.Agreement, error) {
	rows, err := r.findAgreement.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agreements []domain.Agreement

	for rows.Next() {
		var agreement domain.Agreement

		err := rows.Scan(
			&agreement.ID,
			&agreement.Name,
			&agreement.NumberOfExtraPayments,
		)
		if err != nil {
			return nil, err
		}

		log.Printf("FindAgreements returned agreement: ID=%v, Name=%v, ExtraPayments=%v",
			agreement.ID, agreement.Name, agreement.NumberOfExtraPayments)

		agreements = append(agreements, agreement)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return agreements, nil
}
