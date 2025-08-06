package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PayrollIncidentRequest struct {
	PayrollID   string    `json:"payroll_id" binding:"required,uuid"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

func (h Handler) PostPayrollIncident(c *gin.Context) {
	var req PayrollIncidentRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostPayrollIncident request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostPayrollIncident request: %+v", req)

	if err := h.app.CreatePayrollIncident(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create payroll incident: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create payroll incident"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "payroll incident created successfully"})
}
