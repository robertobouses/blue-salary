package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var MigrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env.config file: %v", err)
		}

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPass, dbHost, dbPort, dbName)
		log.Println("Connecting to DB:", dbURL)

		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("failed to connect to DB: %v", err)
		}
		defer db.Close()

		if db.Ping() != nil {
			log.Fatalf("error ping DB: %v", err)
		}

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatalf("failed to create migration driver: %v", err)
		}

		m, err := migrate.NewWithDatabaseInstance("file://cmd/migrations/sql", "postgres", driver)

		if err != nil {
			log.Fatalf("failed to create migrate instance: %v", err)
		}

		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("No new migrations to apply")
			} else {
				log.Fatalf("migration failed: %v", err)
			}
		} else {
			log.Println("Migrations ran successfully")
		}
	},
}
