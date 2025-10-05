-- Create users table
-- This table stores user accounts with authentication and profile information

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(254) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'active', 'suspended')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_status ON users(status);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);
CREATE INDEX IF NOT EXISTS idx_users_updated_at ON users(updated_at);

-- Add trigger to automatically update updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE users IS 'User accounts with authentication and profile information';
COMMENT ON COLUMN users.id IS 'Primary key identifier for the user';
COMMENT ON COLUMN users.email IS 'Unique email address for authentication';
COMMENT ON COLUMN users.password_hash IS 'Bcrypt-hashed password (never store plain text)';
COMMENT ON COLUMN users.name IS 'Display name for the user';
COMMENT ON COLUMN users.status IS 'User account status: pending, active, or suspended';
COMMENT ON COLUMN users.created_at IS 'Timestamp when the user account was created';
COMMENT ON COLUMN users.updated_at IS 'Timestamp when the user account was last updated';