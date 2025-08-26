package company

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	CIF         string `json:"cif"`
	CCC         string `json:"ccc"`
	AgreementID string `json:"agreement_id"`
}

func (h Handler) PostCompany(c *gin.Context) {
	var req CompanyRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostCompany request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostCompany request: %+v", req)

	if err := h.app.CreateCompany(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create company: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "company created successfully"})
}
