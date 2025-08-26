package domain

import "github.com/google/uuid"

type Company struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	CIF         string    `json:"cif"`
	CCC         string    `json:"ccc"`
	AgreementID uuid.UUID `json:"agreement_id"`
}
