SELECT 
    id,
    employee_id,
    description,
    start_date,
    end_date
FROM blues.payroll_incident
WHERE employee_id = $1
  AND start_date <= $3
  AND end_date >= $2;
