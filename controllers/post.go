package controllers

import (
	"example/GoFolderStructure/db/db_connect"
	"example/GoFolderStructure/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) { // creating a new post
	var newPost models.Post
	fmt.Println("in controlless")

	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	insert := `insert into "post" ("PostID","Title","UserName","Description","PostTime") values ($1,$2,$3,$4,$5)`

	d := time.Now().Format("2006-01-02 15:04:05")
	newPost.PostTime = d

	_, err := db.DB.Exec(insert, newPost.PostID, newPost.Title, newPost.UserName, newPost.Description, d)
	if err != nil {
		log.Println("error in query")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in query", "error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func GetPost(c *gin.Context) { // fetching posts from database

	// var posts []post
	// for rows.Next() {
	// 	var PostID int
	// 	var Title string
	// 	var UserName string
	// 	var Description string
	// 	var PostTime string
	// 	err = rows.Scan(&PostID, &Title, &UserName, &Description, &PostTime)
	// 	if err != nil {
	// 		log.Println("scan in get post")
	// 		//c.IndentedJSON(http.StatusConflict, gin.H{"message": "scan in get post", "error": err})
	// 	}
	// 	posts = append(posts, post{PostID: PostID, Title: Title, UserName: UserName, Description: Description, PostTime: PostTime})
	// }
	// fmt.Println(posts)
	//c.IndentedJSON(http.StatusOK, posts)
	rows, err := db.DB.Query("select * from post")
	CheckError(err)
	if err != nil {
		log.Println("not able to fetch data from get post")
	}
	posts := make([]models.Post, 0)

	// var PostID string
	// var Title string
	// var UserName string
	// var Description string
	// var PostTime string
	for rows.Next() {
		p := models.Post{}
		err = rows.Scan(&p.PostID, &p.Title, &p.UserName, &p.Description, &p.PostTime)
		if err != nil {
			log.Println("scan in get post")
		}
		posts = append(posts, p)
	}
	fmt.Println(posts)
	c.IndentedJSON(http.StatusOK, posts)

}
func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}
