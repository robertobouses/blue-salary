package payroll

import (
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) LoadIncidentByEmployeeID(employeeID uuid.UUID, month time.Time) ([]domain.PayrollIncident, error) {

	incidents, err := a.payrollRepo.FindIncidentByEmployeeID(employeeID, month)
	if err != nil {
		return nil, err
	}

	return incidents, nil
}
