package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PayrollMonthRequest struct {
	Month string `json:"month"`
}

func (h *Handler) PostPayrollsPDFByMonth(c *gin.Context) {
	ctx := c.Request.Context()
	var req PayrollMonthRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("Invalid request format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	month, err := time.Parse("2006-01", req.Month)
	if err != nil {
		log.Printf("Invalid month format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month format, use YYYY-MM"})
		return
	}
	outputs, err := h.app.GeneratePayrollsPDFByMonth(ctx, month)
	if err != nil {
		log.Printf("Error generating Payroll PDF: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate PDF"})
		return
	}

	c.JSON(http.StatusOK, outputs)

}
