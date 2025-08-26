package company

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveCompany(ctx context.Context, company domain.Company) error {
	_, err := r.saveCompany.ExecContext(
		ctx,
		company.Name,
		company.Address,
		company.CIF,
		company.CCC,
		company.AgreementID,
	)

	if err != nil {
		log.Print("Error executing SaveCompany statement:", err)
		return err
	}

	return nil
}
