CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create tenants table
CREATE TABLE IF NOT EXISTS tenants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255) UNIQUE,
    plan_id VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    settings JSONB NOT NULL DEFAULT '{}'::JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Insert default tenant
INSERT INTO tenants (id, name, domain, plan_id)
VALUES ('00000000-0000-0000-0000-000000000000', 'Default', 'default.ecomflex.com', 'free')
ON CONFLICT DO NOTHING;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('super_admin', 'influencer', 'public')),
    phone VARCHAR(50),
    referral_code VARCHAR(50) UNIQUE,
    status VARCHAR(50) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'pending_approval', 'rejected')),
    profile_picture VARCHAR(255),
    social_links TEXT[],
    follower_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (email, tenant_id)
);

-- Create index on email for faster lookups
CREATE INDEX idx_users_email ON users(email);

-- Create index on tenant_id for faster lookups
CREATE INDEX idx_users_tenant_id ON users(tenant_id);

-- Create index on role for faster lookups
CREATE INDEX idx_users_role ON users(role);

-- Create index on status for faster lookups
CREATE INDEX idx_users_status ON users(status);

-- Create index on referral_code for faster lookups
CREATE INDEX idx_users_referral_code ON users(referral_code);

-- Insert super admin user (password: admin123)
INSERT INTO users (id, tenant_id, email, password_hash, full_name, role, status)
VALUES (
    '11111111-1111-1111-1111-111111111111',
    '00000000-0000-0000-0000-000000000000',
    'admin@ecomflex.com',
    '$2a$12$sSqDMJ9ZGHPMNGPsE9GX5u.wDnvU7F7OBmjS6XVvLZJP8M0QZ5VZ.',
    'Super Admin',
    'super_admin',
    'active'
)
ON CONFLICT DO NOTHING;

-- Create audit_logs table
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    user_id UUID REFERENCES users(id),
    action VARCHAR(255) NOT NULL,
    entity_type VARCHAR(255) NOT NULL,
    entity_id UUID,
    old_value JSONB,
    new_value JSONB,
    ip_address VARCHAR(45),
    user_agent VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create index on tenant_id for faster lookups
CREATE INDEX idx_audit_logs_tenant_id ON audit_logs(tenant_id);

-- Create index on user_id for faster lookups
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);

-- Create index on created_at for faster lookups
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);