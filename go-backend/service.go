package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseConnection() {

	config := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: "",
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: "user_db",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Datbase connection established!")
}

func GetUsers() ([]UserData, error) {
	var users []UserData
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, fmt.Errorf("error while loading users %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var usr UserData

		if err := rows.Scan(&usr.Id, &usr.Name); err != nil {
			return nil, fmt.Errorf("error while parsing user data %v", err)
		}

		users = append(users, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while reading rows %v", err)
	}

	return users, nil
}

func GetUserById(id int) (UserData, error) {
	var usr UserData
	var err error

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	if err := row.Scan(&usr.Id, &usr.Name); err != nil {
		if err == sql.ErrNoRows {
			return usr, fmt.Errorf("no user with this id %v", id)
		}

		return usr, fmt.Errorf("unexpected error happened %v", err)
	}

	return usr, err
}

func AddUser(user UserData) error {

	result, err := db.Exec("INSERT INTO users (name) values (?)", user.Name)

	if err != nil {
		return fmt.Errorf("error while inserting user %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("error, user not inserted %v", err)
	}

	fmt.Printf("User created with id %d\n", id)

	return err
}
