package agreement

import (
	"log"

	httpagreement "github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

func (a *AppService) LoadAgreements() ([]httpagreement.AgreementResponse, error) {
	agreementsRaw, err := a.agreementRepo.FindAgreements()
	if err != nil {
		log.Printf("[LoadAgreements] Error fetching agreements: %v", err)
		return nil, err
	}

	log.Printf("[LoadAgreements] Found %d agreements", len(agreementsRaw))

	responses := make([]httpagreement.AgreementResponse, 0, len(agreementsRaw))

	for _, agr := range agreementsRaw {
		log.Printf("[LoadAgreements] Processing agreement ID=%v Name=%v", agr.ID, agr.Name)

		categories, err := a.agreementRepo.FindCategoriesByAgreementID(agr.ID)
		if err != nil {
			log.Printf("[LoadAgreements] Error fetching categories for agreement %v: %v", agr.ID, err)
			return nil, err
		}
		log.Printf("[LoadAgreements] Found %d categories for agreement %v", len(categories), agr.ID)

		categoryResponses := make([]httpagreement.CategoryResponse, len(categories))
		for i, cat := range categories {
			log.Printf("[LoadAgreements] Category %v: Name=%v, Level=%v, BaseSalary=%v", cat.ID, cat.Name, cat.Level, cat.BaseSalary)
			categoryResponses[i] = httpagreement.CategoryResponse{
				ID:         cat.ID,
				Name:       cat.Name,
				Level:      cat.Level,
				BaseSalary: cat.BaseSalary,
			}
		}

		complements, err := a.agreementRepo.FindSalaryComplementsByID(agr.ID)
		if err != nil {
			log.Printf("[LoadAgreements] Error fetching salary complements for agreement %v: %v", agr.ID, err)
			return nil, err
		}
		log.Printf("[LoadAgreements] Found %d salary complements for agreement %v", len(complements), agr.ID)

		complementResponses := make([]httpagreement.SalaryComplementResponse, len(complements))
		for i, comp := range complements {
			log.Printf("[LoadAgreements] SalaryComplement %v: Name=%v, Type=%v, Value=%v", comp.ID, comp.Name, comp.Type, comp.Value)
			complementResponses[i] = httpagreement.SalaryComplementResponse{
				ID:    comp.ID,
				Name:  comp.Name,
				Type:  comp.Type,
				Value: comp.Value,
			}
		}

		responses = append(responses, httpagreement.AgreementResponse{
			ID:                    agr.ID,
			Name:                  agr.Name,
			NumberOfExtraPayments: agr.NumberOfExtraPayments,
			Categories:            categoryResponses,
			SalaryComplements:     complementResponses,
		})
	}

	log.Printf("[LoadAgreements] Returning %d agreement responses", len(responses))
	return responses, nil
}
