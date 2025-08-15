package agreement

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/blue-salary/internal/domain"
)

type AgreementUpdateRequest struct {
	ID                    uuid.UUID                 `json:"id"`
	Name                  string                    `json:"name"`
	NumberOfExtraPayments int                       `json:"number_of_extra_payments"`
	Categories            []domain.Category         `json:"categories"`
	SalaryComplements     []domain.SalaryComplement `json:"salary_complements"`
}

func (h Handler) UpdateAgreement(c *gin.Context) {
	var req AgreementUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse UpdateAgreement request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received UpdateAgreement request: %+v", req)

	agreement := domain.Agreement{
		ID:                    req.ID,
		Name:                  req.Name,
		NumberOfExtraPayments: req.NumberOfExtraPayments,
	}
	categories := req.Categories
	salaryComplements := req.SalaryComplements

	if err := h.app.UpdateFullAgreement(
		c.Request.Context(),
		agreement,
		categories,
		salaryComplements,
	); err != nil {
		log.Printf("http: failed to update agreement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update agreement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "agreement updated successfully"})
}
