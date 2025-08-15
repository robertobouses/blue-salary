package employee

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEmployees(c *gin.Context) {
	employees, err := h.app.LoadEmployees()
	if err != nil {
		log.Printf("Error loading employees: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get employees details"})
		return
	}

	var employeeResponses []EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, EmployeeResponse{
			ID:             employee.ID,
			FirstName:      employee.FirstName,
			LastName:       employee.LastName,
			SecondLastName: employee.SecondLastName,
		})
	}
	c.JSON(http.StatusOK, employeeResponses)
}
