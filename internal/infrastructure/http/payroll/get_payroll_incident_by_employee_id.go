package payroll

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayrollIncidentResponse struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (h *Handler) GetPayrollIncidentsByEmployeeID(c *gin.Context) {
	employeeIDString := c.Query("employee_id")
	monthString := c.Query("month")

	employeeID, err := uuid.Parse(employeeIDString)
	if err != nil {
		log.Printf("Invalid employee_id: %s | Error: %v", employeeIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee_id"})
		return
	}

	month, err := time.Parse("2006-01", monthString)
	if err != nil {
		log.Printf("Invalid month format: %s | Error: %v", monthString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month format, expected YYYY-MM"})
		return
	}

	incidents, err := h.app.LoadIncidentByEmployeeID(employeeID, month)
	if err != nil {
		log.Printf("Error loading incidents for employee_id %s: %v", employeeID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get payroll incidents"})
		return
	}

	var response []PayrollIncidentResponse
	for _, i := range incidents {
		response = append(response, PayrollIncidentResponse{
			ID:          i.ID,
			Description: i.Description,
			StartDate:   i.StartDate,
			EndDate:     i.EndDate,
		})
	}

	c.JSON(http.StatusOK, gin.H{"incidents": response})
}
