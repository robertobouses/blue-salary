package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CalculatePayrollRequest struct {
	EmployeeID string    `json:"employee_id"`
	Month      time.Time `json:"month"`
}

func (h Handler) PostCalculatePayrollByEmployeeID(c *gin.Context) {
	var req CalculatePayrollRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: [payroll] invalid request format in PostCalculatePayrollByEmployeeID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: [payroll] received request to calculate payroll for employee_id=%s", req.EmployeeID)
	employeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		log.Printf("usecase: invalid payroll_id format: %v", err)
		return
	}
	payroll, err := h.app.CalculatePayrollByEmployeeID(c.Request.Context(), employeeID, req.Month)
	if err != nil {
		log.Printf("http: [payroll] error calculating payroll for employee_id=%s: %v", req.EmployeeID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not calculate payroll"})
		return
	}

	log.Printf("http: [payroll] successfully calculated payroll for employee_id=%s", req.EmployeeID)
	c.JSON(http.StatusOK, gin.H{
		"message": "payroll calculated successfully",
		"payroll": payroll,
	})
}
