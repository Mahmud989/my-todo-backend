package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"my-todo/models"
	"my-todo/services"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var data models.Todo
	err = json.Unmarshal(value, &data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	todo := services.SaveTodo(data.Content)
	c.IndentedJSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	todos := services.GetTodos()
	c.IndentedJSON(http.StatusOK, todos)
}
