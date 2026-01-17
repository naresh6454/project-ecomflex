-- ecomflex-backend/migrations/000002_create_products_table.up.sql
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    asin VARCHAR(20) NOT NULL,
    founder VARCHAR(255) NOT NULL,
    product_link TEXT NOT NULL,
    details TEXT,
    required_bookings INT NOT NULL DEFAULT 0,
    current_bookings INT NOT NULL DEFAULT 0,
    image_url TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    tenant_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES tenants(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX IF NOT EXISTS products_tenant_id_idx ON products(tenant_id);
CREATE INDEX IF NOT EXISTS products_is_active_idx ON products(is_active);
CREATE INDEX IF NOT EXISTS products_current_bookings_idx ON products(current_bookings);