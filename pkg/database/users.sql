-- name: FindUserById :one
SELECT *
  FROM users
 WHERE id = @id;

-- name: FindUserByUuid :one
SELECT *
  FROM users
 WHERE uuid = @uuid;

-- name: FindUserByDisplayName :one
SELECT *
  FROM users
 WHERE display_name = @display_name;

-- name: InsertUser :execrows
INSERT INTO users
(
	uuid,
	first_name,
	last_name,
	display_name,
	password_hash,
	password_salt,
	avatar_url,
	totp_secret,
	refresh_jti
)
VALUES
(
	@uuid,
	@first_name,
	@last_name,
	@display_name,
	@password_hash,
	@password_salt,
	@avatar_url,
	@totp_secret,
	@refresh_jti
);

-- name: UpdateUser :execrows
UPDATE users
SET    first_name = @first_name,
        last_name = @last_name,
     display_name = @display_name,
    password_hash = @password_hash,
    password_salt = @password_salt,
       avatar_url = @avatar_url,
      totp_secret = @totp_secret,
      refresh_jti = @refresh_jti
WHERE        uuid = @uuid;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE uuid = @uuid;
