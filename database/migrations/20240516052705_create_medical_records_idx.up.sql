CREATE INDEX IF NOT EXISTS idx_mr_patient_identity_number
    ON medical_records (patient_identity_number);
CREATE INDEX IF NOT EXISTS idx_mr_created_by
    ON medical_records (created_by);
CREATE INDEX IF NOT EXISTS idx_mr_created_at_desc
    ON medical_records (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_mr_created_at_asc
    ON medical_records (created_at ASC);