package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	query := "SELECT * FROM users WHERE name = '" + name + "'"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, username)
	}
}
