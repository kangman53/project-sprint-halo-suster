CREATE INDEX IF NOT EXISTS idx_patients_identity_number
    ON patients (identity_number);
CREATE INDEX IF NOT EXISTS idx_patients_name
    ON patients USING HASH(lower(name));
CREATE INDEX IF NOT EXISTS idx_patients_phone_number
    ON patients (phone_number);
CREATE INDEX IF NOT EXISTS idx_patients_created_at_desc
    ON patients (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_patients_created_at_asc
    ON patients (created_at ASC);