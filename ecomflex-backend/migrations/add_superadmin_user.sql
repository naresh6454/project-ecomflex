-- add_superadmin_user.sql
-- This file inserts or updates the superadmin user with hardcoded credentials

-- Insert superadmin user with fixed ID and credentials
INSERT INTO users (
    id, 
    tenant_id, 
    email, 
    password_hash, 
    full_name, 
    role, 
    phone, 
    status, 
    created_at, 
    updated_at
) 
VALUES (
    '00000000-0000-0000-0000-000000000001', 
    '00000000-0000-0000-0000-000000000000', 
    'admin@ecomflex.com', 
    '$2a$10$zRYA8MEXTmXvK9Q5vB7kH.T4AjGYscZGYSfLCpPgEZ.EL69T12hCO', -- bcrypt hash for 'Admin@123'
    'System Administrator', 
    'admin', -- Important: using 'admin' instead of 'super_admin' to match frontend
    '', 
    'active', 
    NOW(), 
    NOW()
) 
ON CONFLICT (email) DO UPDATE 
SET 
    password_hash = '$2a$10$zRYA8MEXTmXvK9Q5vB7kH.T4AjGYscZGYSfLCpPgEZ.EL69T12hCO',
    role = 'admin',
    status = 'active',
    updated_at = NOW();