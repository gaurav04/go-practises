package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Posts struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/rest_api_example")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO posts VALUES ( 0, 'Testing' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	results, err := db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var post Posts

		err = results.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(post.Title)
	}

}
