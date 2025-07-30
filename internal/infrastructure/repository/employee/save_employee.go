package employee

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveEmployee(c context.Context, employee domain.Employee) error {
	_, err := r.saveEmployee.Exec(
		employee.FirstName,
		employee.LastName,
		employee.SecondLastName,
	)

	if err != nil {
		log.Print("Error executing SaveEmployee statement:", err)
		return err
	}

	return nil
}
