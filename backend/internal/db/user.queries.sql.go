// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user.queries.sql

package db

import (
	"context"
	"database/sql"
)

const countAllBlogs = `-- name: CountAllBlogs :one
SELECT
  count(*)
FROM
  blogs
`

func (q *Queries) CountAllBlogs(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllBlogs)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countAllComments = `-- name: CountAllComments :one
SELECT
  count(*)
FROM
  comments
`

func (q *Queries) CountAllComments(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllComments)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countMyBlogs = `-- name: CountMyBlogs :one
SELECT
  count(*)
FROM
  blogs
WHERE
  author_id = ?
`

func (q *Queries) CountMyBlogs(ctx context.Context, authorID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countMyBlogs, authorID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countMyComments = `-- name: CountMyComments :one
SELECT
  count(*)
FROM
  comments
WHERE
  author_id = ?
`

func (q *Queries) CountMyComments(ctx context.Context, authorID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countMyComments, authorID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE
  author_id = ?
  AND title = ?
  AND bid = ?
`

type DeleteBlogParams struct {
	AuthorID int64
	Title    string
	Bid      int64
}

func (q *Queries) DeleteBlog(ctx context.Context, arg DeleteBlogParams) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, arg.AuthorID, arg.Title, arg.Bid)
	return err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE
  cid = ?
  AND author_id = ?
  AND blog_id = ?
`

type DeleteCommentParams struct {
	Cid      int64
	AuthorID int64
	BlogID   int64
}

func (q *Queries) DeleteComment(ctx context.Context, arg DeleteCommentParams) error {
	_, err := q.db.ExecContext(ctx, deleteComment, arg.Cid, arg.AuthorID, arg.BlogID)
	return err
}

const deleteCommentByAId = `-- name: DeleteCommentByAId :exec
DELETE FROM comments
WHERE
  author_id = ?
  AND blog_id = ?
`

type DeleteCommentByAIdParams struct {
	AuthorID int64
	BlogID   int64
}

func (q *Queries) DeleteCommentByAId(ctx context.Context, arg DeleteCommentByAIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteCommentByAId, arg.AuthorID, arg.BlogID)
	return err
}

const getAllBlogs = `-- name: GetAllBlogs :many
SELECT
  bid, title, content, author_id, created_at, updated_at
FROM
  blogs
`

func (q *Queries) GetAllBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getAllBlogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.Bid,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllComments = `-- name: GetAllComments :many
SELECT
  cid, content, created_at, updated_at, author_id, blog_id
FROM
  comments
`

func (q *Queries) GetAllComments(ctx context.Context) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getAllComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Cid,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorID,
			&i.BlogID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorName = `-- name: GetAuthorName :one
SELECT
  username
FROM
  users
WHERE
  uid = ?
`

func (q *Queries) GetAuthorName(ctx context.Context, uid int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getAuthorName, uid)
	var username string
	err := row.Scan(&username)
	return username, err
}

const getBlogByBId = `-- name: GetBlogByBId :one
SELECT
  bid, title, content, author_id, created_at, updated_at
FROM
  blogs
WHERE
  bid = ?
`

func (q *Queries) GetBlogByBId(ctx context.Context, bid int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlogByBId, bid)
	var i Blog
	err := row.Scan(
		&i.Bid,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBlogByTitle = `-- name: GetBlogByTitle :one
SELECT
  bid, title, content, author_id, created_at, updated_at
FROM
  blogs
WHERE
  title = ?
`

func (q *Queries) GetBlogByTitle(ctx context.Context, title string) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlogByTitle, title)
	var i Blog
	err := row.Scan(
		&i.Bid,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCommentsForBlogByBId = `-- name: GetCommentsForBlogByBId :many
SELECT
  bid, title, b.content, b.author_id, b.created_at, b.updated_at, cid, c.content, c.created_at, c.updated_at, c.author_id, blog_id
FROM
  blogs AS b
  JOIN comments AS c ON b.bid = c.blog_id
WHERE
  b.bid = ?
`

type GetCommentsForBlogByBIdRow struct {
	Bid         int64
	Title       string
	Content     string
	AuthorID    int64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Cid         int64
	Content_2   string
	CreatedAt_2 sql.NullTime
	UpdatedAt_2 sql.NullTime
	AuthorID_2  int64
	BlogID      int64
}

func (q *Queries) GetCommentsForBlogByBId(ctx context.Context, bid int64) ([]GetCommentsForBlogByBIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsForBlogByBId, bid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsForBlogByBIdRow
	for rows.Next() {
		var i GetCommentsForBlogByBIdRow
		if err := rows.Scan(
			&i.Bid,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Cid,
			&i.Content_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.AuthorID_2,
			&i.BlogID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentsForUserByAid = `-- name: GetCommentsForUserByAid :many
SELECT
  uid, username, email, password, u.created_at, u.updated_at, cid, content, c.created_at, c.updated_at, author_id, blog_id
FROM
  users AS u
  JOIN comments AS c ON u.uid = c.author_id
WHERE
  u.uid = ?
`

type GetCommentsForUserByAidRow struct {
	Uid         int64
	Username    string
	Email       string
	Password    sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Cid         int64
	Content     string
	CreatedAt_2 sql.NullTime
	UpdatedAt_2 sql.NullTime
	AuthorID    int64
	BlogID      int64
}

func (q *Queries) GetCommentsForUserByAid(ctx context.Context, uid int64) ([]GetCommentsForUserByAidRow, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsForUserByAid, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsForUserByAidRow
	for rows.Next() {
		var i GetCommentsForUserByAidRow
		if err := rows.Scan(
			&i.Uid,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Cid,
			&i.Content,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.AuthorID,
			&i.BlogID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBlogs = `-- name: GetUserBlogs :many
SELECT
  bid, title, content, author_id, created_at, updated_at
FROM
  blogs
WHERE
  author_id = ?
`

func (q *Queries) GetUserBlogs(ctx context.Context, authorID int64) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getUserBlogs, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.Bid,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const newBlog = `-- name: NewBlog :exec
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
  updated_at = CURRENT_TIMESTAMP
`

type NewBlogParams struct {
	Title    string
	Content  string
	AuthorID int64
}

func (q *Queries) NewBlog(ctx context.Context, arg NewBlogParams) error {
	_, err := q.db.ExecContext(ctx, newBlog, arg.Title, arg.Content, arg.AuthorID)
	return err
}

const newComment = `-- name: NewComment :exec
INSERT INTO
  comments (content, author_id, blog_id)
VALUES
  (?, ?, ?) ON DUPLICATE KEY
UPDATE content =
VALUES
  (content),
  updated_at = CURRENT_TIMESTAMP
`

type NewCommentParams struct {
	Content  string
	AuthorID int64
	BlogID   int64
}

func (q *Queries) NewComment(ctx context.Context, arg NewCommentParams) error {
	_, err := q.db.ExecContext(ctx, newComment, arg.Content, arg.AuthorID, arg.BlogID)
	return err
}

const pagination = `-- name: Pagination :many
SELECT
  bid, title, content, author_id, created_at, updated_at
FROM
  blogs
ORDER BY
  created_at DESC
LIMIT
  ?
OFFSET
  ?
`

type PaginationParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) Pagination(ctx context.Context, arg PaginationParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, pagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.Bid,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :exec
UPDATE blogs
SET
  title = ?,
  content = ?,
  author_id = ?
WHERE
  bid = ?
  AND author_id =
VALUES
  (author_id)
`

type UpdateBlogParams struct {
	Title    string
	Content  string
	AuthorID int64
	Bid      int64
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) error {
	_, err := q.db.ExecContext(ctx, updateBlog,
		arg.Title,
		arg.Content,
		arg.AuthorID,
		arg.Bid,
	)
	return err
}

const updateBlogByBId = `-- name: UpdateBlogByBId :exec
UPDATE blogs
SET
  title = ?,
  content = ?,
  author_id = ?
WHERE
  bid = ?
`

type UpdateBlogByBIdParams struct {
	Title    string
	Content  string
	AuthorID int64
	Bid      int64
}

func (q *Queries) UpdateBlogByBId(ctx context.Context, arg UpdateBlogByBIdParams) error {
	_, err := q.db.ExecContext(ctx, updateBlogByBId,
		arg.Title,
		arg.Content,
		arg.AuthorID,
		arg.Bid,
	)
	return err
}

const updateComment = `-- name: UpdateComment :exec
update comments
set
  content = ?
where
  cid = ?
  AND author_id = ?
  AND blog_id = ?
`

type UpdateCommentParams struct {
	Content  string
	Cid      int64
	AuthorID int64
	BlogID   int64
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) error {
	_, err := q.db.ExecContext(ctx, updateComment,
		arg.Content,
		arg.Cid,
		arg.AuthorID,
		arg.BlogID,
	)
	return err
}

const updateCommentbyAId = `-- name: UpdateCommentbyAId :exec
update comments
set
  content = ?
where
  author_id = ?
  AND blog_id = ?
`

type UpdateCommentbyAIdParams struct {
	Content  string
	AuthorID int64
	BlogID   int64
}

func (q *Queries) UpdateCommentbyAId(ctx context.Context, arg UpdateCommentbyAIdParams) error {
	_, err := q.db.ExecContext(ctx, updateCommentbyAId, arg.Content, arg.AuthorID, arg.BlogID)
	return err
}
