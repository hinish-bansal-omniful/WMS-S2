package domain

import "time"

// Hub - Struct representing a hub
type Hub struct {
	ID        int64     `json:"id"` // Corresponds to BIGSERIAL (int64 in Go)
	Name      string    `json:"name"`
	TenantID  int64     `json:"tenant_id"`  // Corresponds to BIGINT (int64 in Go)
	Location  string    `json:"location"`   // Could be a string or use a custom type for the POINT type
	CreatedAt time.Time `json:"created_at"` // Corresponds to TIMESTAMPTZ (string or time.Time in Go)
	CreatedBy int64     `json:"created_by"` // Corresponds to BIGINT (int64 in Go)
	UpdatedAt time.Time `json:"updated_at"` // Corresponds to TIMESTAMPTZ (string or time.Time in Go)
	UpdatedBy int64     `json:"updated_by"` // Corresponds to BIGINT (int64 in Go)
	DeletedAt *string   `json:"deleted_at"` // Nullable, so it's a pointer to string or time.Time
}

type SKU struct {
	ID         int64   `json:"id" db:"id"`                 // Corresponds to the 'id' column (BIGSERIAL)
	SellerID   int64   `json:"seller_id" db:"seller_id"`   // Corresponds to the 'seller_id' column (BIGINT)
	Attributes string  `json:"attributes" db:"attributes"` // Corresponds to the 'attributes' column (JSONB)
	PPU        float64 `json:"ppu" db:"ppu"`               // Corresponds to the 'ppu' column (DECIMAL)
	CreatedAt  string  `json:"created_at" db:"created_at"` // Corresponds to the 'created_at' column (TIMESTAMP)
	UpdatedAt  string  `json:"updated_at" db:"updated_at"` // Corresponds to the 'updated_at' column (TIMESTAMP)
}

func (SKU) TableName() string {
	return "sku"
}

type Inventory struct {
	ID        int64  `json:"id" db:"id"`                 // Corresponds to BIGSERIAL (int64 in Go)
	HubID     int64  `json:"hub_id" db:"hub_id"`         // Foreign key to hubs (BIGINT)
	SKUID     int64  `json:"sku_id" db:"sku_id"`         // Foreign key to sku (BIGINT)
	Quantity  int    `json:"quantity" db:"quantity"`     // Corresponds to INT, default 0
	CreatedAt string `json:"created_at" db:"created_at"` // Corresponds to TIMESTAMP
	UpdatedAt string `json:"updated_at" db:"updated_at"` // Corresponds to TIMESTAMP
}
