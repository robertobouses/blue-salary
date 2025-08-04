package employee

import (
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindEmployees() ([]domain.Employee, error) {
	rows, err := r.findEmployees.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []domain.Employee

	for rows.Next() {
		var employee domain.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.SecondLastName,
		)
		if err != nil {
			return nil, err
		}

		log.Printf("FindEmployees returned employee: ID=%v, FirstName=%v, LastName=%v, SecondLastName=%v",
			employee.ID, employee.FirstName, employee.LastName, employee.SecondLastName)

		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
