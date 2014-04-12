package rdbg

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Store struct {
	DB *sql.DB
}

func (s Store) get(n Name) (*DB, error) {
	_, err := s.DB.Query("SELECT name FROM rdbg WHERE id = $1", n)
	return nil, err
}

func (s Store) add(db *DB) error {
    q := "CREATE DATABASE " + string(db.Name)
    _, err := s.DB.Exec(q)
    if err != nil {
        return err
    }

    _, err = s.DB.Exec("INSERT INTO rdbg (name) VALUES ($1)", string(db.Name))
    if err != nil {
        q = "DROP DATABASE " + string(db.Name)
        _, derr := s.DB.Exec(q)
        if derr != nil {
            log.Fatal(derr)
        }
        return err
    }

    return err
}

func (s Store) rm(db *DB) error {
    _, err := s.DB.Exec("DELETE FROM rdbg WHERE name = $1", string(db.Name))
	return err
}

func (s Store) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rdb := GenerateDB(s)
	j, err := json.Marshal(rdb)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, string(j))
}
