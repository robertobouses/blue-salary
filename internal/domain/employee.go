package domain

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID              uuid.UUID
	FirstName       string
	LastName        string
	SecondLastName  string
	GrossSalary     int
	CategoryID      uuid.UUID
	HireDate        time.Time
	TerminationDate *time.Time
}
