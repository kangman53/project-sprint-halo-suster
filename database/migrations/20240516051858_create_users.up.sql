CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    nip VARCHAR(15),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(6) NOT NULL,
    identity_card_scan_img TEXT,
    is_deleted BOOL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
