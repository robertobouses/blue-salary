INSERT INTO blues.employee(
    first_name,
    last_name,
    second_last_name,
    gross_salary,
    category_id,
    hire_date,
    termination_date
)VALUES(
    $1, $2, $3, $4, $5, $6, $7
);
