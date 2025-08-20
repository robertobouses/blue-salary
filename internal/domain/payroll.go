package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payroll struct {
	ID                     uuid.UUID
	EmployeeID             uuid.UUID
	StartDate              time.Time
	EndDate                time.Time
	ExtraPayment           bool
	BaseSalary             int
	SalaryComplements      []int
	PersonalComplement     int
	ExtraHourPay           int
	MonthlyGrossWithExtras int
	BCCC                   int
	BCCP                   int
	IrpfAmount             int
	IrpfEffectiveRate      int
	SSContributions        int
	NetSalary              int
}

type PayrollSalaryComplement struct {
	ID        uuid.UUID
	PayrollID uuid.UUID
	Name      string
	Type      string
	Value     int
}

type PayrollIncident struct {
	ID          uuid.UUID
	EmployeeID  uuid.UUID
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
