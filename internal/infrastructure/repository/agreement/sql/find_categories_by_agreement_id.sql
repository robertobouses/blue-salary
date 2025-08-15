SELECT
id,
name,
level,
base_salary,
agreement_id
FROM blues.category
WHERE agreement_id=$1
