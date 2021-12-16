package repositories

import (
	"my-todo/models"
	"sync"
	"time"
)

var once sync.Once
var instance TodoRepository
var todoList = make([]models.Todo, 0)
var index uint = 0

type TodoRepository struct {
}

func (c TodoRepository) GetTodos() []models.Todo {
	return todoList
}
func (c TodoRepository) AddTodo(content string) models.Todo {
	index = index+1
	todo := models.Todo{
		Index:     index,
		Content:   content,
		CreatedAt: time.Now(),
	}
	todoList = append(
		todoList,
		todo,
	)
	return todo
}

// GetTodoRepository singleton pattern
func GetTodoRepository() TodoRepository {
	return instance
}
