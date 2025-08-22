package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CalculatePayrollsMonthRequest struct {
	Month time.Time `json:"month"`
}

func (h Handler) PostCalculatePayrollsByMonth(c *gin.Context) {
	var req CalculatePayrollsMonthRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: [payrolls] invalid request format in PostCalculatePayrollsByMonth: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	payrolls, err := h.app.CalculatePayrollsByMonth(c.Request.Context(), req.Month)
	if err != nil {
		log.Printf("http: [payrolls] error calculating payrollss for month=%s: %v", req.Month, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not calculate payrolls"})
		return
	}

	log.Printf("http: [payrolls] successfully calculated payrolls for Month=%s", req.Month)
	c.JSON(http.StatusOK, gin.H{
		"message":  "payrolls calculated successfully",
		"payrolls": payrolls,
	})
}
