-- Create connections table
-- This table stores relationships between users (friend requests and connections)

CREATE TABLE IF NOT EXISTS connections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    requester_id UUID NOT NULL,
    addressee_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'blocked')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- Foreign key constraints
    CONSTRAINT fk_connections_requester_id
        FOREIGN KEY (requester_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_connections_addressee_id
        FOREIGN KEY (addressee_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    -- Ensure users cannot connect to themselves
    CONSTRAINT no_self_connection
        CHECK (requester_id != addressee_id),

    -- Ensure only one connection exists between any two users
    UNIQUE (requester_id, addressee_id)
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_connections_requester_id ON connections(requester_id);
CREATE INDEX IF NOT EXISTS idx_connections_addressee_id ON connections(addressee_id);
CREATE INDEX IF NOT EXISTS idx_connections_status ON connections(status);
CREATE INDEX IF NOT EXISTS idx_connections_requester_status ON connections(requester_id, status);
CREATE INDEX IF NOT EXISTS idx_connections_addressee_status ON connections(addressee_id, status);
CREATE INDEX IF NOT EXISTS idx_connections_created_at ON connections(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_connections_updated_at ON connections(updated_at DESC);

-- Add trigger to automatically update updated_at column
CREATE TRIGGER update_connections_updated_at
    BEFORE UPDATE ON connections
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE connections IS 'Relationships between users (friend requests and connections)';
COMMENT ON COLUMN connections.id IS 'Primary key identifier for the connection';
COMMENT ON COLUMN connections.requester_id IS 'User who initiated the connection request';
COMMENT ON COLUMN connections.addressee_id IS 'User who received the connection request';
COMMENT ON COLUMN connections.status IS 'Connection status: pending, accepted, or blocked';
COMMENT ON COLUMN connections.created_at IS 'Timestamp when the connection was created';
COMMENT ON COLUMN connections.updated_at IS 'Timestamp when the connection was last updated';
COMMENT ON CONSTRAINT no_self_connection IS 'Prevents users from connecting to themselves';
COMMENT ON CONSTRAINT connections_requester_id_addressee_id_key IS 'Ensures only one connection exists between any two users';