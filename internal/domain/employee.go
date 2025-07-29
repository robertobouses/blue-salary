package domain

import "github.com/google/uuid"

type Employee struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	SecondLastName string
}
