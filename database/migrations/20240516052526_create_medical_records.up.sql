CREATE TABLE IF NOT EXISTS medical_records (
    id VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    patient_identity_number BIGINT NOT NULL,
    symptoms TEXT NOT NULL,
    medications TEXT NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_identity_number) REFERENCES patients(identity_number) ON DELETE NO ACTION,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE NO ACTION
);
