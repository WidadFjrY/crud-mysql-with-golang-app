package repository

import (
	"context"
	"crud-mysql-with-golang-app/app/entity"
	"database/sql"
	"errors"
	"strconv"
)

type TodoListRepositoryImpl struct {
	DB *sql.DB
}

func NewTodoListRepository(db *sql.DB) TodoListRepository {
	return &TodoListRepositoryImpl{DB: db}
}

func (repo *TodoListRepositoryImpl) GetAllTodoList(ctx context.Context) ([]entity.TodoList, error) {
	sql := "SELECT id, todo, description, is_finished, created_at FROM todo_list"
	rows, err := repo.DB.QueryContext(ctx, sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todos []entity.TodoList
	for rows.Next() {
		todo := entity.TodoList{}
		rows.Scan(&todo.Id, &todo.Todo, &todo.Description, &todo.IsFinished, &todo.CreatedAt)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *TodoListRepositoryImpl) GetTodoListById(ctx context.Context, id int32) (entity.TodoList, error) {
	sql := "SELECT id, todo, description, is_finished, created_at FROM todo_list WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, sql, id)

	todos := entity.TodoList{}

	if err != nil {
		return todos, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&todos.Id, &todos.Todo, &todos.Description, &todos.IsFinished, &todos.CreatedAt)
		return todos, nil
	} else {
		return todos, errors.New("Todo With Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *TodoListRepositoryImpl) AddTodoList(ctx context.Context, todoList entity.TodoList) (entity.TodoList, error) {
	sql := "INSERT INTO todo_list(todo, description) VALUES (?, ?)"

	result, err := repo.DB.ExecContext(ctx, sql, todoList.Todo, todoList.Description)
	if err != nil {
		return todoList, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return todoList, err
	}

	todoList.Id = int32(id)

	return todoList, err
}

func (repo *TodoListRepositoryImpl) UpdateTodoList(ctx context.Context, todoList entity.TodoList, id int32) (entity.TodoList, error) {
	sql := "UPDATE todo_list SET todo = ?, description = ? WHERE id = ?"

	_, err := repo.DB.ExecContext(ctx, sql, todoList.Todo, todoList.Description, id)
	if err != nil {
		return todoList, err
	}

	return todoList, nil
}

func (repo *TodoListRepositoryImpl) DeleteTodoList(ctx context.Context, id int32) error {
	sql := "DELETE FROM todo_list WHERE id = ?"

	_, err := repo.DB.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodoListRepositoryImpl) IfIdIsExist(ctx context.Context, id int32) (bool, error) {
	sql := "SELECT id FROM todo_list"

	rows, err := repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return false, err
	}

	var idResult []entity.TodoList
	defer rows.Close()
	for rows.Next() {
		result := entity.TodoList{}
		rows.Scan(&result.Id)
		idResult = append(idResult, result)
	}

	for _, result := range idResult {
		if result.Id == id {
			return true, nil
		}
	}

	return false, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
}
