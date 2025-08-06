BEGIN;

CREATE TABLE blues.payroll_incident (
    id UUID PRIMARY KEY uuid_generate_v4(),,
    payroll_id UUID NOT NULL,
    description TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    FOREIGN KEY (payroll_id) REFERENCES payroll(id)
);

COMMIT;
