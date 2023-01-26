package controllers

import (
	"example/GoFolderStructure/db/db_connect"
	"example/GoFolderStructure/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) { // adding comment to a post
	var newComment models.Comments

	if err := c.BindJSON(&newComment); err != nil {
		return
	}

	insert := `insert into "comments" ("CommentID","Comment_post_id","CommentedUser","CommentTime","CommentDescription") values ($1,$2,$3,$4,$5)`
	d := time.Now().Format("2006-01-02 15:04:05")
	newComment.CommentTime = d

	_, err := db.DB.Exec(insert, newComment.CommentID, newComment.Comment_post_id, newComment.CommentedUser, newComment.CommentTime, newComment.CommentDescription)
	CheckError(err)

	c.IndentedJSON(http.StatusCreated, newComment)
}

func GetCommentByPostID(c *gin.Context) { //fetching comment by postid

	// 	// get the id from the url
	id := c.Param("id")
	// fmt.Println(id)
	// get data from database
	// return data
	//d, _ := strconv.Atoi(id)

	rows, err := db.DB.Query(`SELECT * FROM "comments" WHERE "Comment_post_id" = $1`, id)
	CheckError(err)
	// temp := make([]comment, 0)
	// for rows.Next() {
	// 	c := comment{}
	// 	err = rows.Scan(&c.CommentID, &c.PostID, &c.CommentedUser, &c.CommentTime,&c.CommentDescription)
	// 	CheckError(err)
	// 	temp = append(temp, c)
	// }

	fmt.Println(rows)

	var temp models.Comments
	for rows.Next() {
		var commentid string
		var comment_post_id string
		var commenteduser string
		var commenttime string
		var commentdescription string
		err = rows.Scan(&commentid, &comment_post_id, &commenteduser, &commenttime, &commentdescription)
		CheckError(err)
		temp = models.Comments{CommentID: commentid, Comment_post_id: comment_post_id, CommentedUser: commenteduser, CommentTime: commenttime, CommentDescription: commentdescription}
	}
	c.IndentedJSON(http.StatusOK, temp)
}
