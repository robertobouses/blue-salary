// main.go
package main

import (
	"log"

	"github.com/robertobouses/blue-salary/cmd/migrations"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "bluesalary",
		Short: "Blue Salary app CLI",
	}

	rootCmd.AddCommand(migrations.MigrationsCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
