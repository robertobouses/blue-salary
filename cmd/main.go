package main

import (
	"log"

	"github.com/robertobouses/blue-salary/cmd/migrations"
	"github.com/robertobouses/blue-salary/cmd/server"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "bluesalary",
		Short: "Blue Salary app CLI",
	}

	rootCmd.AddCommand(migrations.MigrationsCmd)
	rootCmd.AddCommand(server.ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
