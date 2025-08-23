SELECT 
    id,
    payroll_id,
    name,
    type,
    value
FROM blues.payroll_salary_complement 
WHERE payroll_id = $1;
