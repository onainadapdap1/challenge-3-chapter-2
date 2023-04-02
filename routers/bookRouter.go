package routers

import (
	"sql_api_implementation_2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/books", controllers.CreateBook)

	router.GET("/books", controllers.GetAllBook)

	router.GET("/books/:bookID", controllers.GetBookByID)

	router.PUT("/books/:bookID", controllers.UpdateBookByID)

	router.DELETE("/books/:bookID", controllers.DeleteBookByID)
	return router
}