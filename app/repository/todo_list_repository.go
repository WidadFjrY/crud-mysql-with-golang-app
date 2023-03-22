package repository

import (
	"context"
	"crud-mysql-with-golang-app/app/entity"
)

type TodoListRepository interface {
	GetAllTodoList(ctx context.Context) ([]entity.TodoList, error)
	GetTodoListById(ctx context.Context, id int32) (entity.TodoList, error)
	AddTodoList(ctx context.Context, todoList entity.TodoList) (entity.TodoList, error)
	UpdateTodoList(ctx context.Context, todoList entity.TodoList, id int32) (entity.TodoList, error)
	DeleteTodoList(ctx context.Context, id int32) error
	IfIdIsExist(ctx context.Context, id int32) (bool, error)
}
