package http

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/employee"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/model_145"
)

type Server struct {
	agreement agreement.Handler
	employee  employee.Handler
	model145  model_145.Handler
	engine    *gin.Engine
}

func NewServer(agreement agreement.Handler, employee employee.Handler, model145 model_145.Handler) Server {
	return Server{
		agreement: agreement,
		employee:  employee,
		model145:  model145,
		engine:    gin.Default(),
	}
}

func (s *Server) Run(port string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET, PUT, POST, DELETE, PATCH, OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Accept-Language"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	agreement := s.engine.Group("/agreement")
	agreement.POST("/create", s.agreement.PostAgreement)
	agreement.POST("/category", s.agreement.PostCategory)
	agreement.POST("/complement", s.agreement.PostSalaryComplement)
	agreement.PATCH("/update", s.agreement.UpdateAgreement)
	agreement.GET("/all", s.agreement.GetAgreements)

	employee := s.engine.Group("/employee")
	employee.POST("/create", s.employee.PostEmployee)
	employee.GET("/:id", s.employee.GetEmployeeByID)

	model145 := s.engine.Group("/model145")
	model145.POST("/create", s.model145.PostModel145)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
