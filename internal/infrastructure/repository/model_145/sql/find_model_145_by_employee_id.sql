SELECT 
    id,
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
FROM blues.model_145
WHERE employee_id = $1;
