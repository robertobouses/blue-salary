# blue-salary
Payroll Manager is a backend application written in Go that manages employees and calculates their monthly payslips according to the Spanish labor and tax system.  It uses calcsalary, a custom Go library developed by the same author, to compute payroll breakdowns including gross salary, taxes (IRPF), and social security contributions.



## Setup and Run

### Requirements

- Docker and Docker Compose installed
- Go 1.XX installed

### Start the database
docker-compose up -d

### Run migrations
go run cmd/main.go migrations

### Run migrations
go run cmd/main.go server