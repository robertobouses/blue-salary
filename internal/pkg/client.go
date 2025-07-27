package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

func NewPostgres(c DBConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Pass, c.Host, c.Port, c.Database)
	var db *sql.DB
	var err error

	maxRetries := 4

	for i := 1; i <= maxRetries; i++ {
		db, err = sql.Open("postgres", connString)
		if err != nil {
			log.Printf("attempt %d: error opening connection: %v\n", i, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			log.Println("the connection has the correct credentials")
			return db, nil
		}

		log.Printf("attempt %d: database not ready: %v\n", i, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts: %w", maxRetries, err)
}
