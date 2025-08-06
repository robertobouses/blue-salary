package agreement

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SalaryComplementRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       int    `json:"value"`
	AgreementID string `json:"agreement_id"`
}

func (h Handler) PostSalaryComplement(c *gin.Context) {
	var req SalaryComplementRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostSalaryComplement request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostSalaryComplement request: %+v", req)

	if err := h.app.CreateSalaryComplement(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create salary complement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create salary complement"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "salary complement created successfully"})
}
