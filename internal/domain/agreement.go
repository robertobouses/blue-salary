package domain

import (
	"github.com/google/uuid"
)

type Agreement struct {
	ID                    uuid.UUID
	Name                  string
	NumberOfExtraPayments int
	Categories            []Category
	Complements           []SalaryComplement
}

type Category struct {
	ID          uuid.UUID
	Name        string
	Level       int
	BaseSalary  int
	AgreementID uuid.UUID
}

type SalaryComplement struct {
	ID          uuid.UUID
	Name        string
	Type        string
	Value       int
	AgreementID uuid.UUID
}
