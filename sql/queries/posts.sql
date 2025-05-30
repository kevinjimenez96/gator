-- name: CreatePost :one
INSERT INTO posts (
    id,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1, -- id
    $2, -- updated_at
    $3, -- title
    $4, -- url
    $5, -- description
    $6, -- published_at
    $7  -- feed_id
)
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.id, posts.title, posts.url, posts.description, users.name as username, posts.published_at, feeds.name as feed_name FROM users
JOIN feed_follows ON users.id = feed_follows.user_id
JOIN posts ON feed_follows.feed_id = posts.feed_id
JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE users.name = @username
ORDER BY posts.published_at DESC
LIMIT $1;