-- Create likes table
-- This table stores user likes on posts

CREATE TABLE IF NOT EXISTS likes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    post_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- Foreign key constraints
    CONSTRAINT fk_likes_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_likes_post_id
        FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,

    -- Ensure user can only like a post once
    UNIQUE (user_id, post_id)
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_likes_user_id ON likes(user_id);
CREATE INDEX IF NOT EXISTS idx_likes_post_id ON likes(post_id);
CREATE INDEX IF NOT EXISTS idx_likes_created_at ON likes(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_likes_user_post ON likes(user_id, post_id);

-- Add comments for documentation
COMMENT ON TABLE likes IS 'User likes on posts';
COMMENT ON COLUMN likes.id IS 'Primary key identifier for the like';
COMMENT ON COLUMN likes.user_id IS 'User who gave the like';
COMMENT ON COLUMN likes.post_id IS 'Post that was liked';
COMMENT ON COLUMN likes.created_at IS 'Timestamp when the like was created';
COMMENT ON CONSTRAINT likes_user_id_post_id_key IS 'Ensures a user can only like a post once';