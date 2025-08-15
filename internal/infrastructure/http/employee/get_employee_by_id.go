package employee

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeResponse struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	SecondLastName string    `json:"second_last_name"`
}

func (h *Handler) GetEmployeeByID(c *gin.Context) {
	employeeIDString := c.Param("id")

	employeeID, err := uuid.Parse(employeeIDString)
	if err != nil {
		log.Printf("Invalid employee_id: %s | Error: %v", employeeIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee_id"})
		return
	}

	employee, err := h.app.LoadEmployeeByID(employeeID)
	if err != nil {
		log.Printf("Error loading employee by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get employee details"})
		return
	}

	employeeResponse := EmployeeResponse{
		ID:             employee.ID,
		FirstName:      employee.FirstName,
		LastName:       employee.LastName,
		SecondLastName: employee.SecondLastName,
	}

	c.JSON(http.StatusOK, employeeResponse)
}
