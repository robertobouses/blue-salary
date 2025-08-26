package payroll

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) GeneratePayrollPDFByID(ctx context.Context, payrollID uuid.UUID) (domain.GeneratePayrollPDFOutput, error) {
	payroll, err := a.payrollRepo.FindPayrollByID(ctx, payrollID)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("get payroll: %w", err)
	}
	employee, err := a.employeeRepo.FindEmployeeByID(payroll.EmployeeID)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("get employee: %w", err)
	}

	complements, err := a.payrollRepo.FindSalaryComplementsByPayrollID(ctx, payrollID)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("get salary complements by payroll id: %w", err)
	}
	category, err := a.agreementRepo.FindCategoryByID(employee.CategoryID)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("get category by id: %w", err)
	}
	company, err := a.companyRepo.FindCompanyByAgreementID(ctx, category.AgreementID)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("get salary complements by payroll id: %w", err)
	}

	pdfBytes, err := a.pdfService.RenderPayroll(payroll, complements, employee, company)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("render payroll pdf: %w", err)
	}

	exportDir := "exports/payrolls"
	err = os.MkdirAll(exportDir, os.ModePerm)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("create export dir: %w", err)
	}

	fileName := fmt.Sprintf("%s/payroll_%s.pdf", exportDir, payrollID)

	err = os.WriteFile(fileName, pdfBytes, 0644)
	if err != nil {
		return domain.GeneratePayrollPDFOutput{}, fmt.Errorf("save pdf to disk: %w", err)
	}

	return domain.GeneratePayrollPDFOutput{
		FileName: fmt.Sprintf("payroll_%s.pdf", payrollID),
		Content:  pdfBytes,
	}, nil
}
