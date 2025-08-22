package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayrollResponse struct {
	ID                     uuid.UUID `json:"id"`
	EmployeeID             uuid.UUID `json:"employee_id"`
	StartDate              time.Time `json:"start_date"`
	EndDate                time.Time `json:"end_date"`
	ExtraPayment           bool      `json:"extra_payment"`
	BaseSalary             int       `json:"base_salary"`
	SalaryComplements      []int     `json:"salary_complements"`
	PersonalComplement     int       `json:"personal_complement"`
	ExtraHourPay           int       `json:"extra_hour_pay"`
	MonthlyGrossWithExtras int       `json:"monthly_gross_with_extras"`
	BCCC                   int       `json:"bccc"`
	BCCP                   int       `json:"bccp"`
	IrpfAmount             int       `json:"irpf_amount"`
	IrpfEffectiveRate      int       `json:"irpf_effective_rate"`
	SSContributions        int       `json:"ss_contributions"`
	NetSalary              int       `json:"net_salary"`
}

func (h *Handler) GetPayrollByID(c *gin.Context) {
	payrollIDString := c.Param("id")

	payrollID, err := uuid.Parse(payrollIDString)
	if err != nil {
		log.Printf("Invalid Payroll_id: %s | Error: %v", payrollIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Payroll_id"})
		return
	}

	payroll, err := h.app.LoadPayrollByID(payrollID)
	if err != nil {
		log.Printf("Error loading Payroll by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Payroll details"})
		return
	}

	payrollResponse := PayrollResponse{
		ID:                     payroll.ID,
		EmployeeID:             payroll.EmployeeID,
		StartDate:              payroll.StartDate,
		EndDate:                payroll.EndDate,
		ExtraPayment:           payroll.ExtraPayment,
		BaseSalary:             payroll.BaseSalary,
		SalaryComplements:      payroll.SalaryComplements,
		PersonalComplement:     payroll.PersonalComplement,
		ExtraHourPay:           payroll.ExtraHourPay,
		MonthlyGrossWithExtras: payroll.MonthlyGrossWithExtras,
		BCCC:                   payroll.BCCC,
		BCCP:                   payroll.BCCP,
		IrpfAmount:             payroll.IrpfAmount,
		IrpfEffectiveRate:      payroll.IrpfEffectiveRate,
		SSContributions:        payroll.SSContributions,
		NetSalary:              payroll.NetSalary,
	}

	c.JSON(http.StatusOK, payrollResponse)
}
