CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    identity_number BIGINT UNIQUE,
    phone_number VARCHAR (15) NOT NULL,
    name VARCHAR(30) NOT NULL,
    gender VARCHAR(6) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    identity_card_scan_img TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
