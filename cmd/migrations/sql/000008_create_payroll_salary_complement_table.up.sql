BEGIN;

CREATE TABLE blues.payroll_salary_complement (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payroll_id UUID NOT NULL REFERENCES blues.payroll(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    value NUMERIC NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


COMMIT;
