package service

import "crud-mysql-with-golang-app/app/entity"

type TodoListService interface {
	GetAllTodoListService() []entity.TodoList
	GetTodoListByIdService(id int32) ([]entity.TodoList, error)
	AddTodoListService(todo string, desc string) error
	UpdateTodoListService(todo string, desc string, id int32) error
	DeleteTodoListService(id int32) error
}
