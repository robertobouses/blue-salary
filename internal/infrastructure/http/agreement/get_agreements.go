package agreement

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AgreementResponse struct {
	ID                    uuid.UUID                  `json:"id"`
	Name                  string                     `json:"name"`
	NumberOfExtraPayments int                        `json:"number_of_extra_payments"`
	Categories            []CategoryResponse         `json:"categories"`
	SalaryComplements     []SalaryComplementResponse `json:"salary_complements"`
}

type CategoryResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	BaseSalary int       `json:"base_salary"`
}

type SalaryComplementResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type  string    `json:"type"`
	Value int       `json:"value"`
}

func (h *Handler) GetAgreements(c *gin.Context) {
	agreements, err := h.app.LoadAgreements()
	if err != nil {
		log.Printf("Error loading agreements: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get agreements"})
		return
	}

	var responses []AgreementResponse

	for _, agreement := range agreements {
		var categoryResponses []CategoryResponse
		for _, cat := range agreement.Categories {
			categoryResponses = append(categoryResponses, CategoryResponse{
				ID:         cat.ID,
				Name:       cat.Name,
				Level:      cat.Level,
				BaseSalary: cat.BaseSalary,
			})
		}

		var complementResponses []SalaryComplementResponse
		for _, comp := range agreement.SalaryComplements {
			complementResponses = append(complementResponses, SalaryComplementResponse{
				ID:    comp.ID,
				Name:  comp.Name,
				Type:  comp.Type,
				Value: comp.Value,
			})
		}

		responses = append(responses, AgreementResponse{
			ID:                    agreement.ID,
			Name:                  agreement.Name,
			NumberOfExtraPayments: agreement.NumberOfExtraPayments,
			Categories:            categoryResponses,
			SalaryComplements:     complementResponses,
		})
	}

	c.JSON(http.StatusOK, responses)
}
