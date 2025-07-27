BEGIN;

CREATE TABLE IF NOT EXISTS blues.agreement (
     id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    number_of_extra_payments INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
