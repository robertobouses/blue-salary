BEGIN;

CREATE TABLE blues.model_145 (
    id UUID PRIMARY KEY,
    employee_id UUID NOT NULL REFERENCES blues.employee(id),
    marital_status TEXT NOT NULL,
    has_spouse BOOLEAN NOT NULL,
    spouse_income_below_limit BOOLEAN NOT NULL,
    has_children BOOLEAN NOT NULL,
    children_count INTEGER NOT NULL,
    dependent_children_count INTEGER NOT NULL,
    ascendants_count INTEGER NOT NULL,
    disability_percentage INTEGER NOT NULL,
    is_single_parent_family BOOLEAN NOT NULL,
    mobility_reduced BOOLEAN NOT NULL,
    other_deductions TEXT
);


COMMIT;
