package model_145

import (
	"context"
	"log"

	"github.com/robertobouses/blue-salary/internal/domain"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/model_145"
)

func (a AppService) CreateModel145(ctx context.Context, input model_145.Model145Request) error {
	log.Printf("usecase: creating model145 for employee: %s", input.EmployeeID)

	model145 := domain.Model145{
		EmployeeID:             input.EmployeeID,
		MaritalStatus:          input.MaritalStatus,
		HasSpouse:              input.HasSpouse,
		SpouseIncomeBelowLimit: input.SpouseIncomeBelowLimit,
		HasChildren:            input.HasChildren,
		ChildrenCount:          input.ChildrenCount,
		DependentChildrenCount: input.DependentChildrenCount,
		AscendantsCount:        input.AscendantsCount,
		DisabilityPercentage:   input.DisabilityPercentage,
		IsSingleParentFamily:   input.IsSingleParentFamily,
		MobilityReduced:        input.MobilityReduced,
		OtherDeductions:        input.OtherDeductions,
	}

	if err := a.model145Repo.SaveModel145(ctx, model145); err != nil {
		log.Printf("usecase: failed to save model145: %v", err)
		return err
	}

	log.Println("usecase: model145 saved successfully")
	return nil
}
