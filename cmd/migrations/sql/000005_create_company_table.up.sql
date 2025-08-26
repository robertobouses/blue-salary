BEGIN;

CREATE TABLE IF NOT EXISTS blues.company (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    cif VARCHAR(20) UNIQUE NOT NULL,
    ccc VARCHAR(20) UNIQUE NOT NULL,
    agreement_id UUID NOT NULL,
    CONSTRAINT fk_company_agreement FOREIGN KEY (agreement_id) REFERENCES blues.agreement(id) ON DELETE CASCADE
);

COMMIT;
