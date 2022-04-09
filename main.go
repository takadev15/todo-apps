package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/controllers"
  docs "github.com/takadev15/todo-apps/docs"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  )


func main() {
  router := gin.Default()
  docs.SwaggerInfo.BasePath = "/todos"
  todoRoutes := router.Group("/todos")
  {
    todoRoutes.GET("/test", controllers.TestApi)
    todoRoutes.GET("/")
    todoRoutes.PUT("/:id", controllers.UpdateTodo)
    todoRoutes.DELETE("/:id", controllers.DeleteTodo)
  }
   router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   router.Run(":8080")
}
