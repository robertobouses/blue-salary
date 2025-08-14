SELECT 
    id,
    name,
    level,
    base_salary,
    agreement_id
FROM blues.category
WHERE id = $1
