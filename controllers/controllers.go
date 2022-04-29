package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/models"
)

var (
	todos   = make([]models.Todos, 0, 10)
	counter int
)

type InputModels struct {
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

// GetAll godoc
// @Summary Get details of all TODOS
// @Description Get all TODOS
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todos
// @Router / [get]
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// GetTodo godoc
// @Summary Get details of TODOS by ID
// @Description Get TODOS by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Id"
// @Success 200 {object} models.Todos
// @Router /:id [get]
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

// CreateTodo godoc
// @Summary Create a new TODOS
// @Description Get all TODOS
// @Accept  json
// @Produce  json
// @Param   todo     body    models.Todos     true        "Todo"
// @Success 200 {object} models.Todos
// @Router / [post]
func CreateTodo(c *gin.Context) {
	var (
		inputTodos InputModels
		reqTodos   models.Todos
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
	{
		reqTodos.Id = counter
		reqTodos.Title = inputTodos.Title
		reqTodos.Complete = inputTodos.Complete
	}
	todos = append(todos, reqTodos)
	c.JSON(http.StatusCreated, inputTodos)
}

// UpdateTodo godoc
// @Summary Update or change the todo
// @Description Update or change todo status
// @Accept json
// @Produce json
// @Param todo body models.Todos true "Update todos"
// @Success 200 {object} models.Todos
// @Router /:id [PUT]
func UpdateTodo(c *gin.Context) {
	var (
		result     gin.H
		inputTodos InputModels
	)
	if err := c.ShouldBindJSON(&inputTodos); err != nil {
		result = gin.H{
			"error":   true,
			"message": err,
		}
	}
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	for k, v := range todos {
		if v.Id == id {
			todos[k].Title = inputTodos.Title
			todos[k].Complete = inputTodos.Complete
			result = gin.H{
				"result": todos[k],
			}
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeleteTodo godoc
// @Summary Deleted the todo
// @Description Deleted the todo based on Id
// @Accept json
// @Produce json
// @Param Id path string true "Id"
// @Success 200 {object} models.Todos
// @Router /:id [DELETE]
func DeleteTodo(c *gin.Context) {
	var result gin.H
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	id--
  temp := todos[id]
	todos = append(todos[:id], todos[id+1:]...)
	result = gin.H{
		"deleted todos": temp,
	}
	c.JSON(http.StatusOK, result)
}
