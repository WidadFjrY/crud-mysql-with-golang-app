package view

import (
	"crud-mysql-with-golang-app/app/service"
	"crud-mysql-with-golang-app/util"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ViewTodoList() {
	delayDB()
	fmt.Println("\n----------Todo List App----------")
	userOption()
}

func delayDB() {
	fmt.Println("\nProgram Start")
	time.Sleep(1 * time.Second)
	fmt.Println("Created by Widad FjrY")
	time.Sleep(1 * time.Second)
	fmt.Println("Have a Nice Day 😘")
	time.Sleep(1 * time.Second)
	fmt.Print("\nConnecting Database")

	for i := 1; i <= 5; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}

	db, _ := sql.Open("mysql", "root:Plmokn123@tcp(localhost:3306)/todo_list?parseTime=true")
	_, error := db.Begin()

	if error != nil {
		fmt.Print("\n")
		panic(errors.New("database shutdown"))
	}

	fmt.Println("\nDatabases Connected")
	db.Close()
}

func userOption() {
	fmt.Println("\n1. Get All Todo List")
	fmt.Println("2. Get Todo List By Id")
	fmt.Println("3. Add Todo List")
	fmt.Println("4. Update Todo List")
	fmt.Println("5. Delete Todo List")
	fmt.Println("6. Exit")
	fmt.Print("\nSelect Option : ")
	inputUser := util.InputUserInteger()

	switch inputUser {
	case 1:
		viewGetAllTodoList()
	case 2:
		viewGetTodoListById()
	case 3:
		viewAddTodoList()
	case 4:
		viewUpdateTodoList()
	case 5:
		viewDeleteTodoList()
	case 6:
		fmt.Println("Thanks, Have a Nice Day 😘")
		os.Exit(0)
	default:
		fmt.Println("Option Not Valid")
		userOption()
	}
}

func viewGetAllTodoList() {
	fmt.Println("---------------------------------")
	result := service.GetAllTodoListService()

	for i, todo := range result {
		fmt.Println(i+1, "Id		:", todo.Id, "\n  Todo		:", todo.Todo, "\n  Description	:", todo.Description, "\n  Is Finished	:", todo.IsFinished, "\n  Created At	:", todo.CreatedAt)
		fmt.Println("---------------------------------")
	}
	userOption()
}

func viewGetTodoListById() {
	fmt.Print("Input Id : ")
	result, err := service.GetTodoListByIdService(util.InputUserInteger())

	if err != nil {
		fmt.Println(err)
		viewGetTodoListById()
	}

	todo := result[0]

	fmt.Println("\n---------------------------------")
	fmt.Println("Id		:", todo.Id, "\nTodo		:", todo.Todo, "\nDescription	:", todo.Description, "\nIs Finished	:", todo.IsFinished, "\nCreated At	:", todo.CreatedAt)
	fmt.Println("---------------------------------")

	userOption()
}

func viewAddTodoList() {
	fmt.Println("\n---------------------------------")

	fmt.Print("Type New Todo : ")
	todo, errReadTodo := util.InputUserString()

	if errReadTodo != nil {
		fmt.Println(errReadTodo)
		viewAddTodoList()
	}
	todo = strings.TrimRight(todo, "\n")

	fmt.Print("Add Descripton : ")
	desc, errReadDesc := util.InputUserString()

	if errReadDesc != nil {
		fmt.Println(errReadDesc)
		viewGetAllTodoList()
	}
	desc = strings.TrimRight(desc, "\n")

	err := service.AddTodoListService(todo, desc)
	if err != nil {
		fmt.Println(err)
		viewAddTodoList()
	}

	fmt.Println("Todo List Added...")
	fmt.Println("---------------------------------")

	userOption()
}

func viewUpdateTodoList() {
	fmt.Println("\n---------------------------------")
	fmt.Print("Input id : ")
	id := util.InputUserInteger()

	fmt.Print("Input Change of Todo : ")
	todo, errReadTodo := util.InputUserString()

	if errReadTodo != nil {
		fmt.Println(errReadTodo)
		viewAddTodoList()
	}
	todo = strings.TrimRight(todo, "\n")

	fmt.Print("Input Change of Descripton : ")
	desc, errReadDesc := util.InputUserString()

	if errReadDesc != nil {
		fmt.Println(errReadDesc)
		viewGetAllTodoList()
	}
	desc = strings.TrimRight(desc, "\n")

	err := service.UpdateTodoListService(todo, desc, id)
	if err != nil {
		fmt.Println(err)
		viewAddTodoList()
	}

	fmt.Println("Todo List Change...")
	fmt.Println("---------------------------------")

	userOption()
}

func viewDeleteTodoList() {
	fmt.Println("\n---------------------------------")
	fmt.Print("Input Id : ")

	err := service.DeleteTodoList(util.InputUserInteger())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Todo List Deleted")
	fmt.Println("---------------------------------")

	userOption()
}
