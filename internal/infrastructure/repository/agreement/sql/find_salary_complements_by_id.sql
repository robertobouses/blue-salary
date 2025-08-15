SELECT
id,
name,
type,
value,
agreement_id
FROM blues.salary_complement
WHERE agreement_id=$1
