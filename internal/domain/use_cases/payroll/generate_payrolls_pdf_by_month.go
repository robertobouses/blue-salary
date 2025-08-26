package payroll

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/robertobouses/blue-salary/internal/domain"
)

func (a AppService) GeneratePayrollsPDFByMonth(ctx context.Context, month time.Time) ([]domain.GeneratePayrollPDFOutput, error) {
	payrolls, err := a.payrollRepo.FindPayrollsByMonth(ctx, month)
	if err != nil {
		return nil, fmt.Errorf("get payrolls: %w", err)
	}

	var outputs []domain.GeneratePayrollPDFOutput

	exportDir := "exports/payrolls"
	if err := os.MkdirAll(exportDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("create export dir: %w", err)
	}

	for _, payroll := range payrolls {
		complements, err := a.payrollRepo.FindSalaryComplementsByPayrollID(ctx, payroll.ID)
		if err != nil {
			return nil, fmt.Errorf("get salary complements for payroll %s: %w", payroll.ID, err)
		}

		employee, err := a.employeeRepo.FindEmployeeByID(payroll.EmployeeID)
		if err != nil {
			return nil, fmt.Errorf("render payroll pdf: %w", err)
		}

		pdfBytes, err := a.pdfService.RenderPayroll(payroll, complements, employee)
		if err != nil {
			return nil, fmt.Errorf("render payroll pdf for %s: %w", payroll.ID, err)
		}

		fileName := fmt.Sprintf("%s/payroll_%s.pdf", exportDir, payroll.ID)
		if err := os.WriteFile(fileName, pdfBytes, 0644); err != nil {
			return nil, fmt.Errorf("save pdf to disk: %w", err)
		}

		outputs = append(outputs, domain.GeneratePayrollPDFOutput{
			FileName: fmt.Sprintf("payroll_%s.pdf", payroll.ID),
			Content:  pdfBytes,
		})
	}

	return outputs, nil
}
