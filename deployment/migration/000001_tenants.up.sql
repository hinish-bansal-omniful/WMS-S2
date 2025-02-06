CREATE TABLE IF NOT EXISTS tenants (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by BIGINT,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by BIGINT,
    deleted_at TIMESTAMP DEFAULT NULL
);
