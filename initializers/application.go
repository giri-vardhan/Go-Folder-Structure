package initializers

import (
	db "example/GoFolderStructure/db/db_connect"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApplication() {

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	db.ConnectDB()
	MapUrls()
	router.Run("localhost:8080")
}
