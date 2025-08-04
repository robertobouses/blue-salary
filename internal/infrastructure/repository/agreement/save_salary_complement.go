package agreement

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveSalaryComplement(ctx context.Context, salarycomplement domain.SalaryComplement) error {
	_, err := r.saveAgreement.ExecContext(
		ctx,
		salarycomplement.Name,
		salarycomplement.Type,
		salarycomplement.Value,
		salarycomplement.AgreementID,
	)

	if err != nil {
		log.Print("Error executing SaveSalaryComplement statement:", err)
		return err
	}

	return nil
}
