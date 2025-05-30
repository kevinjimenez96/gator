-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS(
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.user_id as user_id, feed_follows.feed_id as feed_id, users.name as username, feeds.name as feed_name FROM feed_follows
JOIN users ON users.id = feed_follows.user_id AND feed_follows.user_id = $1
JOIN feeds ON feeds.id = feed_follows.feed_id;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
where user_id = $1 AND feed_id = $2;

