package agreement

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgreementRequest struct {
	Name                  string `json:"name"`
	NumberOfExtraPayments int    `json:"number_of_extra_payments"`
}

func (h Handler) PostAgreement(c *gin.Context) {
	var req AgreementRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostAgreement request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostAgreement request: %+v", req)

	if err := h.app.CreateAgreement(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create agreement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create agreement"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "agreement created successfully"})
}
