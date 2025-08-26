SELECT id,
	employee_id,
	start_date,
	end_date,
	extra_payment,
	base_salary,
	personal_complement,
	extra_hour_pay,
	monthly_gross_with_extras,
	bccc,
	bccp,
	irpf_amount,
	irpf_effective_rate,
	ss_contributions,
	net_salary
   FROM blues.payroll
	WHERE start_date >= $1::date
	  AND start_date < ($1::date + interval '1 month')