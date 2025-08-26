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

func (g *Generator) RenderPayroll(payroll domain.Payroll, complements []domain.PayrollSalaryComplement, employee domain.Employee, company domain.Company) ([]byte, error) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("ArialUTF8", "", "internal/infrastructure/pdf/fonts/ARIAL.TTF")
	pdf.AddUTF8Font("ArialUTF8", "B", "internal/infrastructure/pdf/fonts/ARIBLK.TTF")
	pdf.AddUTF8Font("LFAXIDI", "", "internal/infrastructure/pdf/fonts/LFAXDI.TTF")

	pdf.AddPage()
	pdf.SetFont("ArialUTF8", "B", 16)
	pdf.Cell(0, 10, "Payroll Report")
	pdf.Ln(15)
	pdf.SetFont("LFAXIDI", "", 12)

	pdf.Cell(0, 10, fmt.Sprintf("Company: %s", company.Name))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Address: %s", company.Address))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("CIF: %s", company.CIF))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("CCC: %s", company.CCC))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Period: %s to %s", payroll.StartDate.Format("02-01-2006"), payroll.EndDate.Format("02-01-2006")))
	pdf.Ln(12)
	pdf.SetFont("LFAXIDI", "", 12)

	pdf.Cell(0, 10, fmt.Sprintf("Employee ID: %s", payroll.EmployeeID.String()))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Name: %s %s %s", employee.FirstName, employee.LastName, employee.SecondLastName))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Hire Date: %s", employee.HireDate.Format("02-01-2006")))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Period: %s to %s", payroll.StartDate.Format("02-01-2006"), payroll.EndDate.Format("02-01-2006")))
	pdf.Ln(12)

	pdf.SetFont("ArialUTF8", "B", 12)
	pdf.Cell(0, 10, "Salary Details")
	pdf.Ln(10)
	pdf.SetFont("ArialUTF8", "", 12)

	pdf.Cell(0, 8, fmt.Sprintf("Base Salary: €%.2f", float64(payroll.BaseSalary)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Personal Complement: €%.2f", float64(payroll.PersonalComplement)/100))
	pdf.Ln(6)

	if len(complements) > 0 {
		pdf.Cell(0, 8, "Salary Complements:")
		pdf.Ln(8)
		for _, c := range complements {
			pdf.Cell(0, 8, fmt.Sprintf("- %s (%s): €%.2f", c.Name, c.Type, float64(c.Value)/100))
			pdf.Ln(6)
		}
	}

	pdf.Ln(4)
	pdf.Cell(0, 8, fmt.Sprintf("Extra Hour Pay: €%.2f", float64(payroll.ExtraHourPay)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Extra Payment: %t", payroll.ExtraPayment))
	pdf.Ln(12)

	pdf.SetFont("ArialUTF8", "B", 12)
	pdf.Cell(0, 10, "Totals & Deductions")
	pdf.Ln(10)
	pdf.SetFont("ArialUTF8", "", 12)
	pdf.Cell(0, 8, fmt.Sprintf("Gross with Extras: €%.2f", float64(payroll.MonthlyGrossWithExtras)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("BCCC: €%.2f", float64(payroll.BCCC)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("BCCP: €%.2f", float64(payroll.BCCP)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("IRPF: (%.2f%%) €%.2f", float64(payroll.IrpfEffectiveRate)/100, float64(payroll.IrpfAmount)/100))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("SS Contributions: €%.2f", float64(payroll.SSContributions)/100))
	pdf.Ln(12)

	pdf.SetFont("ArialUTF8", "B", 14)
	pdf.Cell(0, 10, fmt.Sprintf("Net Salary: €%.2f", float64(payroll.NetSalary)/100))

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("generate pdf: %w", err)
	}
	return buf.Bytes(), nil
}
