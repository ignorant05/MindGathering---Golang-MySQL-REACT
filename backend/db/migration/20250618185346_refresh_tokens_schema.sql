-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE refresh_tokens (
  tid BIGINT AUTO_INCREMENT PRIMARY KEY,
  token TEXT NOT NULL,
  owner_id BIGINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (owner_id) REFERENCES users (uid) ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS refresh_tokens;
