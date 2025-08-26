package company

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindCompanyByAgreementID(ctx context.Context, agreementID uuid.UUID) (domain.Company, error) {
	row := r.findCompanyByAgreementID.QueryRowContext(ctx, agreementID)

	var company domain.Company
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Address,
		&company.CIF,
		&company.CCC,
		&company.AgreementID,
	)
	if err != nil {
		log.Printf("Repository: error scanning company with AgreementID=%v: %v", agreementID, err)
		return domain.Company{}, err
	}

	log.Printf(
		"Repository: found company [ID=%v, Name=%v, CIF=%v, CCC=%v, AgreementID=%v]",
		company.ID,
		company.Name,
		company.CIF,
		company.CCC,
		company.AgreementID,
	)

	return company, nil
}
