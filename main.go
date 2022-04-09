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
    todoRoutes.GET("/", controllers.GetAll)
    todoRoutes.GET("/:id", controllers.GetTodo)
    todoRoutes.POST("/create", controllers.CreateTodo)
    todoRoutes.PUT("/:id", controllers.UpdateTodo)
    todoRoutes.DELETE("/:id", controllers.DeleteTodo)
  }
   router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   router.Run(":8080")
}
