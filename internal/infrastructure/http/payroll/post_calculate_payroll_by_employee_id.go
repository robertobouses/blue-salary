package payroll

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculatePayrollRequest struct {
	EmployeeID string `json:"employee_id"`
}

func (h Handler) PostCalculatePayrollByEmployeeID(c *gin.Context) {
	var req CalculatePayrollRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostCalculatePayrollByEmployeeID request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostCalculatePayrollByEmployeeID request: %+v", req)

	payroll, err := h.app.CalculatePayrollByEmployeeID(c.Request.Context(), req.EmployeeID)
	if err != nil {
		log.Printf("http: failed to create employee: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully"})
	c.JSON(http.StatusCreated, gin.H{"payroll": payroll})
}
