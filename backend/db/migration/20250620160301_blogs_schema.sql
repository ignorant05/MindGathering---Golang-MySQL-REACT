-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE blogs (
  bid BIGINT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL UNIQUE,
  content TEXT NOT NULL,
  author_id BIGINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS blogs;
