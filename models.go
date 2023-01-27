// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type Comment struct {
	Commentid          int32          `json:"commentid"`
	CommentPostID      sql.NullInt32  `json:"comment_post_id"`
	Commenteduser      sql.NullString `json:"commenteduser"`
	Commenttime        sql.NullString `json:"commenttime"`
	Commentdescription sql.NullString `json:"commentdescription"`
}

type Post struct {
	Postid      int32          `json:"postid"`
	Username    sql.NullString `json:"username"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	Posttime    sql.NullString `json:"posttime"`
}
