package controllers

import (
	"errors"
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
	var reqtodos models.Todos

	err := c.ShouldBindJSON(&reqtodos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	todos = append(todos, reqtodos)
	c.JSON(http.StatusCreated, reqtodos)
}

// swagger
func UpdateTodo(c *gin.Context) {
}

// swagger
func DeleteTodo(c *gin.Context) {
}
