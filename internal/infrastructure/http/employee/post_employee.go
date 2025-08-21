package employee

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EmployeeRequest struct {
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	SecondLastName  string     `json:"second_last_name"`
	GrossSalary     int        `json:"gross_salary"`
	CategoryID      string     `json:"category_id"`
	HireDate        time.Time  `json:"hire_date"`
	TerminationDate *time.Time `json:"termination_date"`
}

func (h Handler) PostEmployee(c *gin.Context) {
	var req EmployeeRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostEmployee request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostEmployee request: %+v", req)

	if err := h.app.CreateEmployee(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create employee: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create employee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully"})
}
