-- name: GetAllBlogs :many
SELECT
  *
FROM
  blogs;

-- name: GetBlogByTitle :one
SELECT
  *
FROM
  blogs
WHERE
  title = ?;

-- name: GetBlogByBId :one
SELECT
  *
FROM
  blogs
WHERE
  bid = ?;

-- name: CountMyBlogs :one 
SELECT
  count(*)
FROM
  blogs
WHERE
  author_id = ?;

-- name: CountAllBlogs :one 
SELECT
  count(*)
FROM
  blogs;

-- name: CountMyComments :one 
SELECT
  count(*)
FROM
  comments
WHERE
  author_id = ?;

-- name: CountAllComments :one 
SELECT
  count(*)
FROM
  comments;

-- name: GetUserBlogs :many
SELECT
  *
FROM
  blogs
WHERE
  author_id = ?;

-- name: GetAllComments :many
SELECT
  *
FROM
  comments;

-- name: NewBlog :exec
INSERT INTO
  blogs (title, content, author_id)
VALUES
  (?, ?, ?) ON DUPLICATE KEY
UPDATE title =
VALUES
  (title),
  content =
VALUES
  (content),
  updated_at = CURRENT_TIMESTAMP;

-- name: UpdateBlogByBId :exec 
UPDATE blogs
SET
  title = ?,
  content = ?,
  author_id = ?
WHERE
  bid = ?;

-- name: UpdateBlog :exec 
UPDATE blogs
SET
  title = ?,
  content = ?,
  author_id = ?
WHERE
  bid = ?
  AND author_id =
VALUES
  (author_id);

-- name: DeleteBlog :exec 
DELETE FROM blogs
WHERE
  author_id = ?
  AND title = ?
  AND bid = ?;

-- name: NewComment :exec 
INSERT INTO
  comments (content, author_id, blog_id)
VALUES
  (?, ?, ?) ON DUPLICATE KEY
UPDATE content =
VALUES
  (content),
  updated_at = CURRENT_TIMESTAMP;

-- name: UpdateCommentbyAId :exec
update comments
set
  content = ?
where
  author_id = ?
  AND blog_id = ?;

-- name: UpdateComment :exec
update comments
set
  content = ?
where
  cid = ?
  AND author_id = ?
  AND blog_id = ?;

-- name: DeleteCommentByAId :exec
DELETE FROM comments
WHERE
  author_id = ?
  AND blog_id = ?;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE
  cid = ?
  AND author_id = ?
  AND blog_id = ?;

-- name: GetCommentsForBlogByBId :many
SELECT
  *
FROM
  blogs AS b
  JOIN comments AS c ON b.bid = c.blog_id
WHERE
  b.bid = ?;

-- name: GetAuthorName :one 
SELECT
  username
FROM
  users
WHERE
  uid = ?;

-- name: GetCommentsForUserByAid :many 
SELECT
  *
FROM
  users AS u
  JOIN comments AS c ON u.uid = c.author_id
WHERE
  u.uid = ?;

-- name: Pagination :many
SELECT
  *
FROM
  blogs
ORDER BY
  created_at DESC
LIMIT
  ?
OFFSET
  ?;
