CREATE TABLE IF NOT EXISTS sellers (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    phone TEXT,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by_id BIGINT,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by_id BIGINT,
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE NULLS NOT DISTINCT (tenant_id, email, deleted_at),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
    );



