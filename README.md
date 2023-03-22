# CRUD MySQL with Golang App

This is a CRUD (Create, Read, Update, Delete) application developed using Golang programming language and MySQL database. This app is implemented using the repository pattern.

## Getting Started

To get started, follow these instructions:

1. Clone this repository using the following command:
   `git clone https://github.com/WidadFjrY/crud-mysql-with-golang-app.git
`
2. Navigate to the project directory:
   `cd crud-mysql-with-golang-app`
3. Install library driver MySQL using the following command: `go get -u https://github.com/go-sql-driver/mysql`
4. Create a database named **todo_list** in your MySQL database.
5. Run the SQL script **db/todo_list.sql** to create the required table.
6. Run the following command to start the application:
   `go run main.go`

## Features

This application provides the following features:

1. Get All Todo List: Retrieves all the todo list items from the database.
2. Get Todo List By ID: Retrieves a todo list item from the database by its ID.
3. Add Todo List: Adds a new todo list item to the database.
4. Update Todo List: Updates an existing todo list item in the database.
5. Delate Todo List: Deletes a todo list item from the database.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/WidadFjrY/crud-mysql-with-golang-app/blob/master/LICENSE) file for details.
