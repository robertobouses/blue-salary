package payroll

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CalculatePersonalComplementRequest struct {
	EmployeeID string `json:"employee_id"`
}

func (h Handler) PostCalculatePersonalComplementByEmployeeID(c *gin.Context) {
	var req CalculatePersonalComplementRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received calculate salary complement request for employee_id=%s", req.EmployeeID)

	employeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		log.Printf("usecase: invalid payroll_id format: %v", err)
		return
	}

	personalComplement, err := h.app.CalculatePersonalComplementByEmployeeID(c.Request.Context(), employeeID)
	if err != nil {
		log.Printf("http: failed to calculate personal complement for employee_id=%s: %v", req.EmployeeID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not calculate personal complement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee_id":        req.EmployeeID,
		"personalComplement": personalComplement,
	})
}
