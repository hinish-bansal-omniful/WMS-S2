CREATE TABLE IF NOT EXISTS sku (
    id  BIGSERIAL PRIMARY KEY,
    seller_id BIGINT NOT NULL,
    attributes JSONB, 
    ppu DECIMAL NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE
);
