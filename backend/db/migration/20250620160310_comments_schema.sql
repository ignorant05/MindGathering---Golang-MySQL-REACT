-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE comments (
  cid BIGINT AUTO_INCREMENT PRIMARY KEY,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  author_id BIGINT NOT NULL,
  blog_id BIGINT NOT NULL,
  FOREIGN KEY (author_id) REFERENCES users (uid),
  FOREIGN KEY (blog_id) REFERENCES blogs (bid)
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS comments;
