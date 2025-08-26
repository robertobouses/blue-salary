package payroll

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayrollIDRequest struct {
	PayrollID string `json:"payroll_id"`
}

func (h *Handler) PostPayrollPDFByID(c *gin.Context) {
	ctx := c.Request.Context()
	var req PayrollIDRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("Invalid request format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	payrollID, err := uuid.Parse(req.PayrollID)
	if err != nil {
		log.Printf("Invalid Payroll_id: %s | Error: %v", req.PayrollID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payroll_id"})
		return
	}

	pdfOutput, err := h.app.GeneratePayrollPDFByID(ctx, payrollID)
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
