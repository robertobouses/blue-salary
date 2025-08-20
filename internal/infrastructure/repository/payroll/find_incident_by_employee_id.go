package payroll

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindIncidentByEmployeeID(employeeID uuid.UUID, month time.Time) ([]domain.PayrollIncident, error) {
	startOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	rows, err := r.findIncidentByEmployeeID.QueryContext(
		context.Background(),
		employeeID,
		startOfMonth,
		endOfMonth,
	)
	if err != nil {
		log.Printf("repository: failed to query payroll incidents: %v", err)
		return nil, err
	}
	defer rows.Close()

	var incidents []domain.PayrollIncident
	for rows.Next() {
		var pi domain.PayrollIncident
		if err := rows.Scan(&pi.ID, &pi.EmployeeID, &pi.Description, &pi.StartDate, &pi.EndDate); err != nil {
			return nil, err
		}
		incidents = append(incidents, pi)
	}

	return incidents, nil
}
