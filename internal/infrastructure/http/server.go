package http

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
)

type Server struct {
	agreement agreement.Handler
	engine    *gin.Engine
}

func NewServer(agreement agreement.Handler) Server {
	return Server{
		agreement: agreement,
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

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
