package payroll

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	calcsalary "github.com/robertobouses/calcsalary/domain"
)

func (a AppService) CalculatePersonalComplementByEmployeeID(ctx context.Context, employeeIDstring string) (int, error) {
	log.Printf("usecase: starting payroll calculation for employee_id=%s", employeeIDstring)

	employeeID, err := uuid.Parse(employeeIDstring)
	if err != nil {
		log.Printf("usecase: invalid payroll_id format: %v", err)
		return 0, err
	}

	employee, err := a.employeeRepo.FindEmployeeByID(employeeID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Employee by ID: %v", err)
		return 0, err
	}

	if employee.CategoryID == uuid.Nil {
		log.Printf("usecase error: employee %s has no category assigned", employee.ID)
		return 0, fmt.Errorf("employee %s has no category assigned", employee.ID)
	}

	log.Printf("usecase: found employee ID=%s | Name=%s %s %s | CategoryID=%s",
		employee.ID, employee.FirstName, employee.LastName, employee.SecondLastName, employee.CategoryID)

	category, err := a.agreementRepo.FindCategoryByID(employee.CategoryID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Category by ID: %v", err)
		return 0, err
	}
	if category.AgreementID == uuid.Nil {
		log.Printf("usecase error: category %s has no agreement assigned", category.ID)
		return 0, fmt.Errorf("category %s has no agreement assigned", category.ID)
	}

	log.Printf("usecase: found category ID=%s | BaseSalary=%d | AgreementID=%s",
		category.ID, category.BaseSalary, category.AgreementID)

	agreement, err := a.agreementRepo.FindAgreementByID(category.AgreementID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Agreement by ID: %v", err)
		return 0, err
	}

	log.Printf("usecase: found agreement ID=%s | ExtraPayments=%d",
		agreement.ID, agreement.NumberOfExtraPayments)

	salaryComplements, err := a.agreementRepo.FindSalaryComplementsByID(category.AgreementID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Salary Complements by ID: %v", err)
		return 0, err
	}

	log.Printf("usecase: found %d salary complements", len(salaryComplements))

	var salaryComplementsValues []int
	for _, sc := range salaryComplements {
		salaryComplementsValues = append(salaryComplementsValues, sc.Value)
	}
	if len(salaryComplementsValues) == 0 {
		log.Printf("usecase warning: employee %s has no salary complements", employee.ID)
	}

	model145, err := a.model145Repo.FindModel145ByEmployeeID(ctx, employeeID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Model145 by Employee ID: %v", err)
		return 0, err
	}

	log.Printf("usecase: found model145 | Children=%d | Disability=%d%% | MobilityReduced=%v | Ascendants=%d | DisabledAscendants=%v",
		model145.ChildrenCount, model145.DisabilityPercentage, model145.MobilityReduced, model145.AscendantsCount, model145.HasDisabledAscendants)

	if category.BaseSalary <= 0 {
		log.Printf("usecase error: category %s has BaseSalary <= 0, cannot calculate payroll", category.ID)
		return 0, fmt.Errorf("category %s has invalid BaseSalary: %d", category.ID, category.BaseSalary)
	}

	payrollInput := calcsalary.PayrollInput{
		BaseSalary:            category.BaseSalary,
		SalaryComplements:     salaryComplementsValues,
		PersonalComplement:    0,
		GrossSalary:           employee.GrossSalary,
		NumberOfExtraPayments: agreement.NumberOfExtraPayments,

		NumberOfExtraHours: 0,
		ExtraHourRate:      0,
		MonthlyHours:       0,

		NumberOfChildren: model145.ChildrenCount,

		HasDisability:       model145.DisabilityPercentage > 0,
		HasSevereDisability: model145.DisabilityPercentage >= 65,
		NeedsAssistance:     model145.MobilityReduced,

		HasAscendantsOver65:   model145.AscendantsCount > 0,
		HasDisabledAscendants: model145.HasDisabledAscendants,
	}
	log.Printf("usecase: calling calcsalary with input: %+v", payrollInput)
	log.Printf("usecase: constructed payrollInput=%+v", payrollInput)

	monthlyPersonalComplement := calcsalary.MonthlyPersonalComplement(payrollInput)
	payrollInput.PersonalComplement = monthlyPersonalComplement
	log.Printf("usecase: calculated personal complement=%d", monthlyPersonalComplement)

	return monthlyPersonalComplement, nil

}
