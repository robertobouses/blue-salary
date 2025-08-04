SELECT
id,
name,
number_of_extra_payments,
FROM blues.agreement
WHERE agreement_id=$1
