package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error connecting: %s", err)
		panic(err)
	}

	b, err := ioutil.ReadFile("query.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	_, err = db.Exec(string(b))
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
