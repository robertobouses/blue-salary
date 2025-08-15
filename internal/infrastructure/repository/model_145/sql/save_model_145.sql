INSERT INTO blues.model_145 (
    employee_id,
    marital_status,
    has_spouse,
    spouse_income_below_limit,
    has_children,
    children_count,
    dependent_children_count,
    ascendants_count,
    disability_percentage,
    is_single_parent_family,
    mobility_reduced,
    has_disabled_ascendants,
    other_deductions
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
);
