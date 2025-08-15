BEGIN;

CREATE TABLE IF NOT EXISTS blues.employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    second_last_name TEXT,
    gross_salary INTEGER NOT NULL,
    category_id UUID REFERENCES blues.category(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
