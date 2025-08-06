package agreement

import (
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a *AppService) LoadAgreements() ([]*domain.Agreement, error) {
	agreementsRaw, err := a.agreementRepo.FindAgreements()
	if err != nil {
		return nil, err
	}

	agreements := make([]*domain.Agreement, len(agreementsRaw))
	for i := range agreementsRaw {
		agreement := agreementsRaw[i]
		categories, err := a.agreementRepo.FindCategoriesByID(agreement.ID)
		if err != nil {
			return nil, err
		}

		complements, err := a.agreementRepo.FindSalaryComplementsByID(agreement.ID)
		if err != nil {
			return nil, err
		}

		agreement.Categories = categories
		agreement.SalaryComplements = complements

		agreements[i] = &agreement
	}

	return agreements, nil
}
