package models

type Post struct {
	PostID      int    `json:"postid"` // primary key
	Title       string `json:"title"`
	UserName    string `json:"username"`
	Description string `json:"description"`
	PostTime    string `json:"posttime"`
}

type Comments struct {
	CommentID          string `json:"commentid"`       // primary key
	Comment_post_id    string `json:"comment_post_id"` //foreign key  referencing to primary key PostID in post table
	CommentedUser      string `json:"commenteduser"`
	CommentTime        string `json:"commenttime"`
	CommentDescription string `json:"commentdescription"`
}
