-- +goose Up
-- +goose StatementBegin
CREATE TABLE recommendation_history (
    id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    recommended_jobs JSON NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_recommendation_history_user 
        FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- Add indexes for faster lookups
CREATE INDEX idx_recommendation_history_user_id ON recommendation_history(user_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_recommendation_history_user_id ON recommendation_history;
DROP TABLE IF EXISTS recommendation_history;
-- +goose StatementEnd
