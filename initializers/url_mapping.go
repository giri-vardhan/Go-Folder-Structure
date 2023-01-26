package initializers

import (
	"example/GoFolderStructure/controllers"
	"fmt"
)

func MapUrls() {
	fmt.Println("api started")

	router.GET("/posts", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.GET("/comments/:id", controllers.GetCommentByPostID)
	router.POST("/comments", controllers.CreateComment)

}
