SELECT
id,
first_name,
last_name,
second_last_name,
gross_salary,
category_id,
hire_date,
termination_date
FROM blues.employee
WHERE id=$1
