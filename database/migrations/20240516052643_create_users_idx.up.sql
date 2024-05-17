CREATE UNIQUE INDEX IF NOT EXISTS unique_nip 
    ON users(nip) WHERE is_deleted = false;
CREATE INDEX IF NOT EXISTS index_users_id
    ON users (id);
CREATE INDEX IF NOT EXISTS index_users_nip
    ON users (nip);
CREATE INDEX IF NOT EXISTS index_users_name
    ON users USING HASH(lower(name));
CREATE INDEX IF NOT EXISTS index_users_role
    ON users (role);
CREATE INDEX IF NOT EXISTS index_users_created_at_desc
    ON users(created_at DESC);
CREATE INDEX IF NOT EXISTS index_users_created_at_asc
    ON users(created_at ASC);
