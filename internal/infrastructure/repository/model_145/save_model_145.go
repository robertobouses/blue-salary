package model_145

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (r *Repository) SaveModel145(ctx context.Context, model145 domain.Model145) error {
	_, err := r.saveModel145.ExecContext(
		ctx,
		model145.EmployeeID,
		model145.MaritalStatus,
		model145.HasSpouse,
		model145.SpouseIncomeBelowLimit,
		model145.HasChildren,
		model145.ChildrenCount,
		model145.DependentChildrenCount,
		model145.AscendantsCount,
		model145.DisabilityPercentage,
		model145.IsSingleParentFamily,
		model145.MobilityReduced,
		model145.HasDisabledAscendants,
		model145.OtherDeductions,
	)

	if err != nil {
		log.Printf("Error executing SaveModel145 statement: %v", err)
		return err
	}

	return nil
}
