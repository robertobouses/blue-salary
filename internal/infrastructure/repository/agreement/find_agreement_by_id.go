package agreement

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindAgreementByID(agreementID uuid.UUID) (domain.Agreement, error) {
	var agreement domain.Agreement

	row := r.findAgreementByID.QueryRow(agreementID)
	err := row.Scan(
		&agreement.ID,
		&agreement.Name,
		&agreement.NumberOfExtraPayments,
	)
	if err != nil {
		return domain.Agreement{}, err
	}

	log.Printf("FindAgreementByID returned agreement: ID=%v, Name=%v, ExtraPayments=%v",
		agreement.ID, agreement.Name, agreement.NumberOfExtraPayments)

	return agreement, nil
}
