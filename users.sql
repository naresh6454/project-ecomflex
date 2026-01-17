CREATE TABLE IF NOT EXISTS users (
       id UUID PRIMARY KEY,
       tenant_id UUID NOT NULL,
       email VARCHAR(255) NOT NULL UNIQUE,
       password_hash VARCHAR(255) NOT NULL,
       full_name VARCHAR(255) NOT NULL,
       phone VARCHAR(50),
       role VARCHAR(20) NOT NULL,
       status VARCHAR(20) NOT NULL DEFAULT 'active',
       referral_code VARCHAR(50),
       social_links TEXT[],
       follower_count INTEGER DEFAULT 0,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP,
       last_login TIMESTAMP
   );