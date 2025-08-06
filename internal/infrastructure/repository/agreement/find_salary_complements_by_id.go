package agreement

import (
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindSalaryComplementsByID(agreementID uuid.UUID) ([]domain.SalaryComplement, error) {
	rows, err := r.findSalaryComplementByID.Query(agreementID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salaryComplements []domain.SalaryComplement

	for rows.Next() {
		var salaryComplement domain.SalaryComplement

		err := rows.Scan(
			&salaryComplement.ID,
			&salaryComplement.Name,
			&salaryComplement.Type,
			&salaryComplement.Value,
			&salaryComplement.AgreementID,
		)
		if err != nil {
			return nil, err
		}

		log.Printf("FindsalaryComplement returned salaryComplement: ID=%v, Name=%v, Type=%v, Value=%v, AgreementID=%v",
			salaryComplement.ID, salaryComplement.Name, salaryComplement.Type, salaryComplement.Value, salaryComplement.AgreementID)

		salaryComplements = append(salaryComplements, salaryComplement)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return salaryComplements, nil
}
