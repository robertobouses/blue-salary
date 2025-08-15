BEGIN;

CREATE TABLE IF NOT EXISTS blues.category (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    level INTEGER NOT NULL,
    base_salary INTEGER NOT NULL,
    agreement_id UUID NOT NULL REFERENCES blues.agreement(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
