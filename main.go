package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/takadev15/todo-apps/controllers"
	docs "github.com/takadev15/todo-apps/docs"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/todos"
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("/test", controllers.TestApi)
		todoRoutes.GET("/getall", controllers.GetAll)
		todoRoutes.GET("/getbyid/:id", controllers.GetTodo)
		todoRoutes.POST("/create", controllers.CreateTodo)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
