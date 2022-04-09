package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/models"
)

var todos = make([]models.Todos, 0, 10)

// test json
func TestApi(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "Message": "testing",
  })
}

// swagger
func GetAll(c *gin.Context) {
  result := gin.H{
    "result" : todos,
  }
  c.JSON(http.StatusOK, result)
}

// swagger
func GetById(c *gin.Context) {
}

// swagger
func CreateTodo(c *gin.Context) {
  var(
    result gin.H
    todo models.Todos
    )
  if err := c.ShouldBindJSON(&todo); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
  }
  todos = append(todos, todo)
  result = gin.H{
    "Added" : todo,
  }
  c.JSON(http.StatusOK, result)
}

// swagger
func UpdateTodo(c *gin.Context) {
  var result gin.H
  inputId := c.Param("id")
  id, _ := strconv.Atoi(inputId)
  for _, v := range(todos) {
    if v.Id == id {
      result = gin.H{
        "result" : v,
      }
    }
  }
  c.JSON(http.StatusOK, result)
}

// swagger
func DeleteTodo(c *gin.Context) {
}

