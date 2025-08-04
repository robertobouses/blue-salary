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
		log.Printf("Invalid employeeID: %s | Error: %v", employeeIDString, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match_id"})
		return
	}

	if employeeID == uuid.Nil {
		log.Println("Missing 'employeeID' query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "employeeID query param is required"})
		return
	}

	employee, err := h.app.LoadEmployeeByID(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get match details"})
		return
	}

	var employeeResponse EmployeeResponse
	employeeResponse.ID = employee.ID
	employeeResponse.FirstName = employee.FirstName
	employeeResponse.LastName = employee.LastName
	employeeResponse.SecondLastName = employee.SecondLastName

	c.JSON(http.StatusOK, employeeResponse)
}
