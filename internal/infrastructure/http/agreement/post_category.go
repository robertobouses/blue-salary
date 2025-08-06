package agreement

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryRequest struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	BaseSalary  int    `json:"base_salary"`
	AgreementID string `json:"agreement_id"`
}

func (h Handler) PostCategory(c *gin.Context) {
	var req CategoryRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostCategory request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostCategory request: %+v", req)

	if err := h.app.CreateCategory(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category created successfully"})
}
