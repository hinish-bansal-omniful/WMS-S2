CREATE TABLE if not EXISTS hubs (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name text NOT NULL,
    tenant_id BIGINT NOT NULL,
    location POINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);
