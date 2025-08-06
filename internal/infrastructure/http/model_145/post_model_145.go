package model_145

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Model145Request struct {
	EmployeeID             string `json:"employee_id"`
	MaritalStatus          string `json:"marital_status"`
	HasSpouse              bool   `json:"has_spouse"`
	SpouseIncomeBelowLimit bool   `json:"spouse_income_below_limit"`
	HasChildren            bool   `json:"has_children"`
	ChildrenCount          int    `json:"children_count"`
	DependentChildrenCount int    `json:"dependent_children_count"`
	AscendantsCount        int    `json:"ascendants_count"`
	DisabilityPercentage   int    `json:"disability_percentage"`
	IsSingleParentFamily   bool   `json:"is_single_parent_family"`
	MobilityReduced        bool   `json:"mobility_reduced"`
	OtherDeductions        string `json:"other_deductions"`
}

func (h Handler) PostModel145(c *gin.Context) {
	var req Model145Request

	if err := c.BindJSON(&req); err != nil {
		log.Printf("http: failed to parse PostModel145 request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	log.Printf("http: received PostModel145 request: %+v", req)

	if err := h.app.CreateModel145(c.Request.Context(), req); err != nil {
		log.Printf("http: failed to create model145: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create model145"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "model145 created successfully"})
}
