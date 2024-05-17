CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(100) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    identity_number VARCHAR(20) UNIQUE,
    phone_number VARCHAR (20) NOT NULL,
    name VARCHAR(40) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    identity_card_scan_img TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
