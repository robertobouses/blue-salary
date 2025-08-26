BEGIN;

CREATE TABLE blues.model_145 (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
    has_disabled_ascendants BOOLEAN NOT NULL,
    other_deductions TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


COMMIT;
