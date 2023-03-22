package service

import (
	"context"
	"crud-mysql-with-golang-app/app/entity"
	"crud-mysql-with-golang-app/app/repository"
	"crud-mysql-with-golang-app/db"
	"errors"
	"strconv"
)

func GetAllTodoListService() []entity.TodoList {
	todoList := repository.NewTodoListRepository(db.GetConnection())
	ctx := context.Background()
	var todo []entity.TodoList

	todos, err := todoList.GetAllTodoList(ctx)

	if err != nil {
		panic(err)
	}

	for _, result := range todos {
		todosResult := entity.TodoList{
			Id:          result.Id,
			Todo:        result.Todo,
			Description: result.Description,
			IsFinished:  result.IsFinished,
			CreatedAt:   result.CreatedAt,
		}
		todo = append(todo, todosResult)
	}

	return todo
}

func GetTodoListByIdService(id int32) ([]entity.TodoList, error) {
	todoList := repository.NewTodoListRepository(db.GetConnection())
	ctx := context.Background()
	var todo []entity.TodoList

	result, err := todoList.GetTodoListById(ctx, id)
	if err != nil {
		return todo, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}

	todoResult := entity.TodoList{
		Id:          result.Id,
		Todo:        result.Todo,
		Description: result.Description,
		IsFinished:  result.IsFinished,
		CreatedAt:   result.CreatedAt,
	}

	return append(todo, todoResult), nil
}

func AddTodoListService(todo string, desc string) error {
	todoList := repository.NewTodoListRepository(db.GetConnection())
	ctx := context.Background()

	newTodoList := entity.TodoList{
		Todo:        todo,
		Description: desc,
	}

	_, err := todoList.AddTodoList(ctx, newTodoList)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodoListService(todo string, desc string, id int32) error {
	todoList := repository.NewTodoListRepository(db.GetConnection())
	ctx := context.Background()

	newTodoList := entity.TodoList{
		Todo:        todo,
		Description: desc,
	}

	_, err := todoList.UpdateTodoList(ctx, newTodoList, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoList(id int32) error {
	todoList := repository.NewTodoListRepository(db.GetConnection())
	ctx := context.Background()

	ifExist, err := todoList.IfIdIsExist(ctx, id)
	if err != nil {
		return err
	}

	if ifExist {
		err := todoList.DeleteTodoList(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}
