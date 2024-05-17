CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(100) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    nip VARCHAR(50),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    identity_card_scan_img TEXT,
    is_deleted BOOL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
