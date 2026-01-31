-- Filename: migrations/000003_add_indexes.up.sql
-- Creates indexes for better query performance

-- Index on username for faster lookup
CREATE INDEX idx_users_username ON users(username);
-- Index on user_id for faster joins and filtering
CREATE INDEX idx_posts_user_id ON posts(user_id);
