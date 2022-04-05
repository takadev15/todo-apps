package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// test json
func TestApi(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "Message": "testing",
  })
}

// swagger
func GetAll(c *gin.Context) {
}

// swagger
func GetById(c *gin.Context) {
}

// swagger
func CreateTodo(c *gin.Context) {
}

// swagger
func UpdateTodo(c *gin.Context) {
}

// swagger
func DeleteTodo(c *gin.Context) {
}
