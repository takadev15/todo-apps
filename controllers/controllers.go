package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/models"
)

var todos = make([]models.Todos, 0, 10)

var counter int

type InputModels struct {
  Title string
  Description string
}

// test json
func TestApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "testing",
	})
}

// swagger
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// swagger
func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	singleTodo, err := GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusOK, singleTodo)
}
func GetById(id int) (*models.Todos, error) {

	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("data not found")
}

// swagger
func CreateTodo(c *gin.Context) {
	var (
    inputTodos InputModels
    reqTodos models.Todos
    )

	err := c.ShouldBindJSON(&inputTodos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
  counter++
  reqTodos.Id = counter
  reqTodos.Title = inputTodos.Title
  reqTodos.Description = inputTodos.Description
	todos = append(todos, reqTodos)
	c.JSON(http.StatusCreated, inputTodos)
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

