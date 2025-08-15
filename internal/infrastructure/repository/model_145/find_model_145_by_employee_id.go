package model_145

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) FindModel145ByEmployeeID(ctx context.Context, employeeID uuid.UUID) (domain.Model145, error) {
	log.Printf("repository: searching for model145 for employee ID %s", employeeID)

	row := r.findModel145ByEmployeeID.QueryRowContext(ctx, employeeID)
	var m domain.Model145
	err := row.Scan(
		&m.ID,
		&m.EmployeeID,
		&m.MaritalStatus,
		&m.HasSpouse,
		&m.SpouseIncomeBelowLimit,
		&m.HasChildren,
		&m.ChildrenCount,
		&m.DependentChildrenCount,
		&m.AscendantsCount,
		&m.DisabilityPercentage,
		&m.IsSingleParentFamily,
		&m.MobilityReduced,
		&m.HasDisabledAscendants,
		&m.OtherDeductions,
	)
	if err != nil {
		log.Printf("repository: failed to find model145 for employee ID %s: %v", employeeID, err)
		return domain.Model145{}, fmt.Errorf("could not find model145 for employee ID %s: %w", employeeID, err)
	}

	log.Printf("repository: found model145: %+v", m)
	return m, nil
}
