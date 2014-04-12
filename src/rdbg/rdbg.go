package rdbg

import (
	"bytes"
	"log"
	"math/rand"
	"time"
)

func GenerateDB(s Store) *DB {
	n := generateName()
	var db *DB

	for db == nil {
		var err error
		db, err = createDB(n)
		if err != nil {
			log.Print(err)
			continue
		}

		err = s.add(db)
		if err != nil {
			db = nil
			log.Print(err)
			continue
		}
	}

	return db
}

func generateName() Name {
	var b bytes.Buffer
	for i := 0; i < 10; i++ {
		c := string(rand.Intn(26) + 97)
		b.WriteString(c)
	}
	return Name(b.String())
}

func createDB(n Name) (*DB, error) {
	db := &DB{Name: n, Created: time.Now()}
	return db, nil
}
