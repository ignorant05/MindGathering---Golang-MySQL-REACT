-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE images (
  iid BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(63) NOT NULL,
  data LONGBLOB NOT NULL,
  user_id BIGINT UNIQUE,
  FOREIGN KEY (user_id) REFERENCES users (uid) ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS images;
