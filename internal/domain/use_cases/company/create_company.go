package company

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/company"
)

func (a AppService) CreateCompany(ctx context.Context, input company.CompanyRequest) error {
	log.Printf("usecase: creating company with name: %s, address: %s, cif: %s, ccc: %s and agreementID: %s", input.Name, input.Address, input.CIF, input.CCC, input.AgreementID)

	agreementID, err := uuid.Parse(input.AgreementID)
	if err != nil {
		log.Printf("usecase: invalid agreementID format: %v", err)
		return err
	}

	company := domain.Company{
		Name:        input.Name,
		Address:     input.Address,
		CIF:         input.CIF,
		CCC:         input.CCC,
		AgreementID: agreementID,
	}

	if err := a.companyRepo.SaveCompany(ctx, company); err != nil {
		log.Printf("usecase: failed to save company: %v", err)
		return err
	}

	log.Println("usecase: company saved successfully")
	return nil
}
