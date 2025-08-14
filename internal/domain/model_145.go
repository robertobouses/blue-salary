package domain

import (
	"github.com/google/uuid"
)

type Model145 struct {
	ID                     uuid.UUID
	EmployeeID             uuid.UUID
	MaritalStatus          string
	HasSpouse              bool
	SpouseIncomeBelowLimit bool
	HasChildren            bool
	ChildrenCount          int
	DependentChildrenCount int
	AscendantsCount        int
	DisabilityPercentage   int
	IsSingleParentFamily   bool
	MobilityReduced        bool
	HasDisabledAscendants  bool
	OtherDeductions        string
}
