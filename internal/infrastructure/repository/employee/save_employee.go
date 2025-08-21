package employee

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveEmployee(ctx context.Context, employee domain.Employee) error {
	_, err := r.saveEmployee.ExecContext(
		ctx,
		employee.FirstName,
		employee.LastName,
		employee.SecondLastName,
		employee.GrossSalary,
		employee.CategoryID,
		employee.HireDate,
		employee.TerminationDate,
	)

	if err != nil {
		log.Print("Error executing SaveEmployee statement:", err)
		return err
	}

	return nil
}
