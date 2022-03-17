package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)


//Запросы


func home_page(w http.ResponseWriter, r *http.Request) {
	database, _ := sql.Open("sqlite3", "./persons.db") //Для открытия соединения с базой данных используем функцию sql.Open()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS persons (Id INTEGER PRIMARY KEY,email TEXT,phone TEXT,firstName TEXT,lastName TEXT)")
	statement.Exec()
	rows, _ := database.Query("SELECT Id,email,phone,firstName,lastName FROM persons")
}

