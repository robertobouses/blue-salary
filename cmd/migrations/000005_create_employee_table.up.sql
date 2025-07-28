BEGIN;

CREATE TABLE blues.employee (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    second_last_name TEXT
);

COMMIT;
