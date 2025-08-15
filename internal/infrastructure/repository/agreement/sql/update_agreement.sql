UPDATE blues.agreement
SET
    name = $1,
    number_of_extra_payments = $2
WHERE id = $3;
