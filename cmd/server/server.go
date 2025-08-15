package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	appAgreement "github.com/robertobouses/blue-salary/internal/domain/use_cases/agreement"
	appEmployee "github.com/robertobouses/blue-salary/internal/domain/use_cases/employee"
	appModel145 "github.com/robertobouses/blue-salary/internal/domain/use_cases/model_145"
	appPayroll "github.com/robertobouses/blue-salary/internal/domain/use_cases/payroll"
	httpServer "github.com/robertobouses/blue-salary/internal/infrastructure/http"
	handlerAgreement "github.com/robertobouses/blue-salary/internal/infrastructure/http/agreement"
	handlerEmployee "github.com/robertobouses/blue-salary/internal/infrastructure/http/employee"
	handlerModel145 "github.com/robertobouses/blue-salary/internal/infrastructure/http/model_145"
	handlerPayroll "github.com/robertobouses/blue-salary/internal/infrastructure/http/payroll"
	repositoryAgreement "github.com/robertobouses/blue-salary/internal/infrastructure/repository/agreement"
	repositoryEmployee "github.com/robertobouses/blue-salary/internal/infrastructure/repository/employee"
	repositoryModel145 "github.com/robertobouses/blue-salary/internal/infrastructure/repository/model_145"
	repositoryPayroll "github.com/robertobouses/blue-salary/internal/infrastructure/repository/payroll"
	internalPostgres "github.com/robertobouses/blue-salary/internal/pkg/postgres"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to get env:", err)
		}

		requiredEnv := []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}
		for _, env := range requiredEnv {
			if os.Getenv(env) == "" {
				log.Fatalf("missing required environment variable: %s", env)
			}
		}

		db, err := internalPostgres.NewPostgres(internalPostgres.DBConfig{
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		})
		if err != nil {
			log.Fatal("failed to connect to database:", err)
		}

		agreementRepo, err := repositoryAgreement.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init agreement repository:", err)
		}
		employeeRepo, err := repositoryEmployee.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init employee repository:", err)
		}
		model145Repo, err := repositoryModel145.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init model 145 repository:", err)
		}
		payrollRepo, err := repositoryPayroll.NewRepository(db)
		if err != nil {
			log.Fatal("failde to init payroll repository:", err)
		}

		agreementApp := appAgreement.NewApp(agreementRepo)
		employeeApp := appEmployee.NewApp(employeeRepo)
		model145App := appModel145.NewApp(model145Repo)
		payrollApp := appPayroll.NewApp(payrollRepo, employeeRepo, agreementRepo, model145Repo)

		agreementHandler := handlerAgreement.NewHandler(&agreementApp)
		employeeHandler := handlerEmployee.NewHandler(employeeApp)
		model145Handler := handlerModel145.NewHandler(model145App)
		payrollHandler := handlerPayroll.NewHandler(payrollApp)

		s := httpServer.NewServer(agreementHandler, employeeHandler, model145Handler, payrollHandler)

		if err := s.Run("8080"); err != nil {
			log.Fatal("server failed:", err)
		}
	},
}
