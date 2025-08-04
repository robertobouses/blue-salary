package employee

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindEmployeeByID(employeeID uuid.UUID) (domain.Employee, error) {
	row := r.findEmployeeByID.QueryRow(employeeID)
	var employee domain.Employee
	err := row.Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.SecondLastName,
	)
	log.Printf("FindEmployeeByID returned employee: ID=%v, FirstName=%v, LastName=%v, SecondLastName=%v", employee.ID, employee.FirstName, employee.LastName, employee.SecondLastName)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}
