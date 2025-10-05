-- Create posts table
-- This table stores user posts with text content and metadata

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    content TEXT NOT NULL CHECK (length(content) >= 1 AND length(content) <= 2000),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- Foreign key constraint
    CONSTRAINT fk_posts_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_posts_user_created ON posts(user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_posts_updated_at ON posts(updated_at DESC);

-- Add trigger to automatically update updated_at column
CREATE TRIGGER update_posts_updated_at
    BEFORE UPDATE ON posts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE posts IS 'User posts with text content and metadata';
COMMENT ON COLUMN posts.id IS 'Primary key identifier for the post';
COMMENT ON COLUMN posts.user_id IS 'Foreign key referencing the user who created the post';
COMMENT ON COLUMN posts.content IS 'Text content of the post (1-2000 characters)';
COMMENT ON COLUMN posts.created_at IS 'Timestamp when the post was created';
COMMENT ON COLUMN posts.updated_at IS 'Timestamp when the post was last updated';