BEGIN;

CREATE TABLE blues.employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    second_last_name TEXT
);

COMMIT;
