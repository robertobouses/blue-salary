package payroll

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
	calcsalary "github.com/robertobouses/calcsalary/domain"
)

const (
	MonthlyDays = 30
)

func (a AppService) CalculatePayrollByEmployeeID(ctx context.Context, employeeIDstring string, month time.Time) (domain.Payroll, error) {
	log.Printf("usecase: starting payroll calculation for employee_id=%s", employeeIDstring)

	employeeID, err := uuid.Parse(employeeIDstring)
	if err != nil {
		return domain.Payroll{}, fmt.Errorf("invalid employee ID format: %v", err)
	}

	employee, err := a.employeeRepo.FindEmployeeByID(employeeID)
	if err != nil {
		return domain.Payroll{}, fmt.Errorf("employee not found: %v", err)
	}

	if employee.CategoryID == uuid.Nil {
		return domain.Payroll{}, fmt.Errorf("employee %s has no category assigned", employee.ID)
	}

	category, err := a.agreementRepo.FindCategoryByID(employee.CategoryID)
	if err != nil {
		return domain.Payroll{}, err
	}

	if category.BaseSalary <= 0 {
		return domain.Payroll{}, fmt.Errorf("category %s has invalid BaseSalary: %d", category.ID, category.BaseSalary)
	}

	agreement, err := a.agreementRepo.FindAgreementByID(category.AgreementID)
	if err != nil {
		return domain.Payroll{}, err
	}

	salaryComplements, err := a.agreementRepo.FindSalaryComplementsByID(category.AgreementID)
	if err != nil {
		return domain.Payroll{}, err
	}

	var salaryComplementsValues []int
	for _, sc := range salaryComplements {
		salaryComplementsValues = append(salaryComplementsValues, sc.Value)
	}

	model145, err := a.model145Repo.FindModel145ByEmployeeID(ctx, employeeID)
	if err != nil {
		return domain.Payroll{}, err
	}

	incidents, err := a.LoadIncidentByEmployeeID(employeeID, month)
	if err != nil {
		return domain.Payroll{}, err
	}

	daysOff := 0
	for _, inc := range incidents {
		days := int(inc.EndDate.Sub(inc.StartDate).Hours()/24) + 1
		daysOff += days
	}

	reductionFactor := 1.0
	if daysOff > 0 {
		if daysOff >= MonthlyDays {
			reductionFactor = 0
		} else {
			reductionFactor = float64(MonthlyDays-daysOff) / float64(MonthlyDays)
		}
	}

	// reducedBase := int(float64(category.BaseSalary) * reductionFactor)
	// reducedComplements := applyFactor(salaryComplementsValues, reductionFactor)

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

	log.Println("input before calcsalary.GeneratePayrollOutput(payrollInput)", payrollInput)
	// monthlyPersonalComplement := calcsalary.MonthlyPersonalComplement(payrollInput)
	// payrollInput.PersonalComplement = int(float64(monthlyPersonalComplement) * reductionFactor)

	output := calcsalary.GeneratePayrollOutput(payrollInput)

	log.Println("output after calcsalary.GeneratePayrollOutput(payrollOutput)", output)

	log.Println("output calcsalary.GeneratePayrollOutput(payrollOutput), output.PersonalComplement", output.PersonalComplement)
	log.Println("output calcsalary.GeneratePayrollOutput(payrollOutput), output.PersonalComplement with reduction", int(float64(output.PersonalComplement)*reductionFactor))

	payroll := domain.Payroll{
		EmployeeID:             employeeID,
		BaseSalary:             int(float64(output.BaseSalary) * reductionFactor),
		SalaryComplements:      applyFactor(output.SalaryComplements, reductionFactor),
		PersonalComplement:     int(float64(output.PersonalComplement) * reductionFactor),
		ExtraHourPay:           output.ExtraHoursPay,
		MonthlyGrossWithExtras: int(float64(output.MonthlyGrossWithExtras) * reductionFactor),
		BCCC:                   int(float64(output.BaseBCCC) * reductionFactor),
		BCCP:                   int(float64(output.BaseBCCP) * reductionFactor),
		IrpfAmount:             output.IrpfAmount,
		IrpfEffectiveRate:      output.IrpfEffectiveRate,
		SSContributions:        output.SSContributions.TotalWorker,
		NetSalary:              output.NetSalary,
	}

	log.Println("output calcsalary.GeneratePayrollOutput(payrollOutput), payroll.PersonalComplement with reduction", payroll.PersonalComplement)

	if err := a.payrollRepo.SavePayroll(ctx, &payroll); err != nil {
		return domain.Payroll{}, err
	}

	log.Println("reductionFactor before SavePayrollSalaryComplement", reductionFactor)

	for _, sc := range salaryComplements {
		payrollSC := domain.PayrollSalaryComplement{
			PayrollID: payroll.ID,
			Name:      sc.Name,
			Type:      sc.Type,
			Value:     int(float64(sc.Value) * reductionFactor),
		}
		if err := a.payrollRepo.SavePayrollSalaryComplement(ctx, payrollSC); err != nil {
			return domain.Payroll{}, err
		}
	}

	return payroll, nil
}

func applyFactor(values []int, factor float64) []int {
	res := make([]int, len(values))
	for i, v := range values {
		res[i] = int(float64(v) * factor)
	}
	return res
}
