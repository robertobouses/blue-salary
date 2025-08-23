package payroll

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetPayrollPDF(c *gin.Context) {
	ctx := c.Request.Context()
	payrollIDString := c.Param("id")

	payrollID, err := uuid.Parse(payrollIDString)
	if err != nil {
		log.Printf("Invalid Payroll_id: %s | Error: %v", payrollIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Payroll_id"})
		return
	}

	pdfOutput, err := h.app.GeneratePayrollPDF(ctx, payrollID)
	if err != nil {
		log.Printf("Error generating Payroll PDF: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate PDF"})
		return
	}

	c.Data(
		http.StatusOK,
		"application/pdf",
		pdfOutput.Content,
	)
}
