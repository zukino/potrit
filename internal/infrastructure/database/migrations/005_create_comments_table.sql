-- Create comments table
-- This table stores user comments on posts

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    post_id UUID NOT NULL,
    content TEXT NOT NULL CHECK (length(content) >= 1 AND length(content) <= 1000),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- Foreign key constraints
    CONSTRAINT fk_comments_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comments_post_id
        FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_comments_post_created ON comments(post_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_comments_user_post ON comments(user_id, post_id);
CREATE INDEX IF NOT EXISTS idx_comments_created_at ON comments(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_comments_updated_at ON comments(updated_at DESC);

-- Add trigger to automatically update updated_at column
CREATE TRIGGER update_comments_updated_at
    BEFORE UPDATE ON comments
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE comments IS 'User comments on posts';
COMMENT ON COLUMN comments.id IS 'Primary key identifier for the comment';
COMMENT ON COLUMN comments.user_id IS 'User who wrote the comment';
COMMENT ON COLUMN comments.post_id IS 'Post being commented on';
COMMENT ON COLUMN comments.content IS 'Text content of the comment (1-1000 characters)';
COMMENT ON COLUMN comments.created_at IS 'Timestamp when the comment was created';
COMMENT ON COLUMN comments.updated_at IS 'Timestamp when the comment was last updated';