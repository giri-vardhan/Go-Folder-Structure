-- name: InsertPost :exec
INSERT INTO post (Title,UserName,Description,PostTime) values($1,$2,$3,$4);

-- name: QueryGetAllPost :many
SELECT * FROM post;

-- name: QueryGetCommentById :many
SELECT * FROM "comments" where Comment_post_id = $1; 

-- name: QueryGetPostById :one
SELECT * FROM "post" where PostID = $1;

-- name: InsertComment :exec
insert into "comments" (Comment_post_id,CommentedUser,CommentTime,CommentDescription) values ($1,$2,$3,$4);
