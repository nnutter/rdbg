package main

import (
	"database/sql"
	//	"fmt"
	"log"
	"net/http"
	"rdbg"

	_ "github.com/lib/pq"
	//	"time"
)

func main() {
	db, err := sql.Open("postgres", "user=nnutter dbname=nnutter sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//	id := 1
	//	rows, err := db.Query("SELECT name FROM rdbg WHERE id = $1", id)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	for rows.Next() {
	//		var name string
	//		if err := rows.Scan(&name); err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Printf("%s\n", name)
	//	}
	//	if err := rows.Err(); err != nil {
	//		log.Fatal(err)
	//	}

	s := rdbg.Store{DB: db}
	http.ListenAndServe("localhost:4000", s)
}
