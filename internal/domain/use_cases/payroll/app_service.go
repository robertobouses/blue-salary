package payroll

import (
	"context"

	"github.com/robertobouses/blue-salary/internal/domain"
)

type PayrollRepository interface {
	SavePayrollIncident(c context.Context, payrollIncident domain.PayrollIncident) error
}

func NewApp(payrollRepository PayrollRepository) AppService {
	return AppService{
		payrollRepo: payrollRepository,
	}
}

type AppService struct {
	payrollRepo PayrollRepository
}
