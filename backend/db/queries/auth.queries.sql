-- name: AllUsers :many 
SELECT
  *
FROM
  users;

-- name: AllTokens :many
SELECT
  *
FROM
  refresh_tokens;

-- name: Register :exec
INSERT INTO
  users (username, email, password)
VALUES
  (?, ?, ?);

-- name: Login :one
SELECT
  *
FROM
  users
where
  username = ?
  AND email = ?;

-- name: Logout :exec 
DELETE FROM refresh_tokens
WHERE
  owner_id = ?;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE
  uid = ?;

-- name: UpdateToken :exec 
UPDATE refresh_tokens
SET
  token = ?
WHERE
  owner_id = ?;

-- name: UpdateUserPic :exec
INSERT INTO
  images (name, type, data, user_id)
VALUES
  (?, ?, ?, ?) ON DUPLICATE KEY
UPDATE name =
VALUES
  (name),
  type =
VALUES
  (type),
  data =
VALUES
  (data);

-- name: GetUserProfilePic :one 
SELECT
  name,
  type,
  data
FROM
  images
WHERE
  user_id = ?;

-- name: UpdateUser :exec
UPDATE users
SET
  username = ?,
  email = ?,
  password = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE
  uid = ?;

-- name: CreateToken :exec
INSERT INTO
  refresh_tokens (token, owner_id)
VALUES
  (?, ?) ON DUPLICATE KEY
UPDATE token =
VALUES
  (token),
  created_at = CURRENT_TIMESTAMP;

-- name: GetUserById :one 
SELECT
  *
FROM
  users
WHERE
  uid = ?;

-- name: GetUserByUsername :one
SELECT
  *
FROM
  users
WHERE
  username = ?;

-- name: GetUserByTid :one 
SELECT
  *
FROM
  users AS u
  JOIN refresh_tokens AS rf ON u.uid = rf.owner_id
WHERE
  tid = ?;

-- name: GetTokenByUid :one 
SELECT
  *
FROM
  refresh_tokens
WHERE
  owner_id = ?;
