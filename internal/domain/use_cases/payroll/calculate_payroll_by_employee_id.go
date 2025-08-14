package payroll

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	calcsalary "github.com/robertobouses/calcsalary/domain"
)

func (a AppService) CalculatePayrollByEmployeeID(ctx context.Context, employeeIDstring string) (domain.Payroll, error) {

	employeeID, err := uuid.Parse(employeeIDstring)
	if err != nil {
		log.Printf("usecase: invalid payroll_id format: %v", err)
		return domain.Payroll{}, err
	}

	employee, err := a.employeeRepo.FindEmployeeByID(employeeID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Employee by ID: %v", err)
		return domain.Payroll{}, err
	}

	category, err := a.agreementRepo.FindCategoryByID(employee.CategoryID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Category by ID: %v", err)
		return domain.Payroll{}, err
	}

	agreement, err := a.agreementRepo.FindAgreementByID(category.AgreementID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Agreement by ID: %v", err)
		return domain.Payroll{}, err
	}

	salaryComplements, err := a.agreementRepo.FindSalaryComplementsByID(category.AgreementID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Salary Complements by ID: %v", err)
		return domain.Payroll{}, err
	}

	var salaryComplementsValues []int
	for _, sc := range salaryComplements {
		salaryComplementsValues = append(salaryComplementsValues, sc.Value)
	}

	model145, err := a.model145Repo.FindModel145ByEmployeeID(ctx, employeeID)
	if err != nil {
		log.Printf("usecase error: CalculatePayrollByEmployeeID, error Find Model145 by Employee ID: %v", err)
		return domain.Payroll{}, err
	}

	payrollInput := calcsalary.PayrollInput{
		BaseSalary:            category.BaseSalary,
		SalaryComplements:     salaryComplementsValues,
		PersonalComplement:    0,
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

	monthlyPersonalComplement := calcsalary.MonthlyPersonalComplement(payrollInput)
	payrollInput.PersonalComplement = monthlyPersonalComplement

	output := calcsalary.GeneratePayrollOutput(payrollInput)

	payroll := domain.Payroll{
		EmployeeID:             employeeID,
		BaseSalary:             output.BaseSalary,
		SalaryComplements:      output.SalaryComplements,
		PersonalComplement:     output.PersonalComplement,
		ExtraHourPay:           output.ExtraHoursPay,
		MonthlyGrossWithExtras: output.MonthlyGrossWithExtras,
		BCCC:                   output.BaseBCCC,
		BCCP:                   output.BaseBCCP,
		IrpfAmount:             output.IrpfAmount,
		IrpfEffectiveRate:      output.IrpfEffectiveRate,
		SSContributions:        output.SSContributions.TotalWorker,
		NetSalary:              output.NetSalary,
	}

	fmt.Printf("Net Salary: %d\n", payroll.NetSalary)
	fmt.Printf("IRPF: %d\n", payroll.IrpfAmount)

	if err := a.payrollRepo.SavePayroll(ctx, payroll); err != nil {
		log.Printf("usecase: failed to save payroll incident: %v", err)
		return domain.Payroll{}, err
	}

	return payroll, nil
}
