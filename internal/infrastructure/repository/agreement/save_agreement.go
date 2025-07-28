package agreement

import (
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveAgreementQuery(agreement domain.Agreement) error {
	_, err := r.saveAgreement.Exec(
		agreement.Name,
		agreement.NumberOfExtraPayments,
	)

	if err != nil {
		log.Print("Error executing SaveAgreement statement:", err)
		return err
	}

	return nil
}
