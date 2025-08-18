SELECT
id,
first_name,
last_name,
second_last_name,
gross_salary,
category_id
FROM blues.employee
WHERE id=$1
