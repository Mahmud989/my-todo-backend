package main

import (
	"my-todo/services"
	"strconv"
	"testing"
)

func TestAddTodo(t *testing.T) {
	todos := services.GetTodos()
	if len(todos) != 0 {
		t.Error("excepted count 0")
	}
	for i := 1; i <= 1000; i++ {
		content := "test todo" + strconv.Itoa(i)
		todo := services.SaveTodo(content)
		if todo.Content != content {
			t.Error("excepted " + content)
		}
		todos = services.GetTodos()
		if len(todos) != i {
			t.Error("excepted count " + strconv.Itoa(i))
		}
	}
}

func TestIndex(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		content := "test todo" + strconv.Itoa(i)
		todo := services.SaveTodo(content)
		if todo.Content != content || todo.Index != uint(i) {
			t.Errorf("excepted content: %v & index: %v", content, i)
		}
		todos := services.GetTodos()
		if len(todos) != i {
			t.Error("excepted count " + strconv.Itoa(i))
		}
	}
}
