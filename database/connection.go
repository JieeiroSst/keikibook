package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func handleDataBase() {
	fmt.Println("Database")

	//open databse mysql
	db, err := sql.Open("mysql", "root: @tcp(127.0.0.1:3306)/keikibook")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//ping to databse
	//Check connection to database
	err = db.Ping()
	if err != nil {
		log.Fatal("connection to failed data")
	}

	fmt.Println("connected to mysql database")

	insert, err := db.Query("insert into user value")
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	fmt.Println("sucesfully inserted into user tables")
}
