-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE users (
  uid BIGINT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email TEXT NOT NULL,
  password VARCHAR(127),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS users;
