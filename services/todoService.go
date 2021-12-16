package services

import (
	"my-todo/models"
	"my-todo/repositories"
	"sync"
)

var once sync.Once

var repository = repositories.GetTodoRepository()


func SaveTodo(content string) models.Todo {

	return repository.AddTodo(content)
}
func GetTodos() []models.Todo {
	return repository.GetTodos()
}
