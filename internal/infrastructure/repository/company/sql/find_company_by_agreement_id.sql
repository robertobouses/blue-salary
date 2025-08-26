SELECT 
    id,
    name,
    address,
    cif,
    ccc,
    agreement_id
FROM blues.company
WHERE agreement_id = $1;
