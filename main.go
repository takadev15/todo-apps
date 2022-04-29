package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/takadev15/todo-apps/controllers"
	docs "github.com/takadev15/todo-apps/docs"
)

// @title           ToDo Apps API
// @version         1.0
// @description     Simple apps for managing todos

// @contact.name   Daffa Haryadi
// @contact.email  daffahr15@protonmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /todos

func main() {
  router := gin.Default()
  docs.SwaggerInfo.BasePath = "/todos"
  todoRoutes := router.Group("/todos")
  {
    todoRoutes.GET("/", controllers.GetAll)
    todoRoutes.GET("/:id", controllers.GetTodo)
    todoRoutes.POST("/", controllers.CreateTodo)
    todoRoutes.PUT("/:id", controllers.UpdateTodo)
    todoRoutes.DELETE("/:id", controllers.DeleteTodo)
  }
   router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   router.Run(":8080")
}
