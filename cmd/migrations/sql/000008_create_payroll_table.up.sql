BEGIN;

CREATE TABLE blues.payroll (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID NOT NULL REFERENCES blues.employee(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    extra_payment BOOLEAN NOT NULL,
    base_salary INTEGER NOT NULL,
    personal_complement INTEGER NOT NULL,
    extra_hour_pay INTEGER NOT NULL,
    monthly_gross_with_extras INTEGER NOT NULL,
    bccc INTEGER NOT NULL,
    bccp INTEGER NOT NULL,
    irpf_amount INTEGER NOT NULL,
    irpf_effective_rate INTEGER NOT NULL,
    ss_contributions INTEGER NOT NULL,
    net_salary INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;
