INSERT INTO blues.payroll (
employee_id,
start_date,
end_date,
extra_payment,
base_salary,
salary_complements,
personal_complement,
extra_hour_pay,
monthly_gross_with_extras,
bccc,
bccp,
irpf_amount,
irpf_effective_rate,
ss_contributions,
netsalary
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
