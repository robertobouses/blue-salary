package pdf

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) RenderPayroll(payroll domain.Payroll, complements []domain.PayrollSalaryComplement) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Payroll Report")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Employee ID: %s", payroll.EmployeeID.String()))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Period: %s to %s", payroll.StartDate.Format("02-01-2006"), payroll.EndDate.Format("02-01-2006")))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Salary Details")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, fmt.Sprintf("Base Salary: €%d", payroll.BaseSalary))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Personal Complement: €%d", payroll.PersonalComplement))
	pdf.Ln(6)

	if len(complements) > 0 {
		pdf.Cell(0, 8, "Salary Complements:")
		pdf.Ln(8)
		for _, c := range complements {
			pdf.Cell(0, 8, fmt.Sprintf("- %s (%s): €%d", c.Name, c.Type, c.Value))
			pdf.Ln(6)
		}
	}

	pdf.Ln(4)
	pdf.Cell(0, 8, fmt.Sprintf("Extra Hour Pay: €%d", payroll.ExtraHourPay))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Extra Payment: %t", payroll.ExtraPayment))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Totals & Deductions")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, fmt.Sprintf("Gross with Extras: €%d", payroll.MonthlyGrossWithExtras))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("BCCC: €%d", payroll.BCCC))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("BCCP: €%d", payroll.BCCP))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("IRPF: %d%%  (€%d)", payroll.IrpfEffectiveRate, payroll.IrpfAmount))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("SS Contributions: €%d", payroll.SSContributions))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, fmt.Sprintf("Net Salary: €%d", payroll.NetSalary))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("generate pdf: %w", err)
	}
	return buf.Bytes(), nil
}
