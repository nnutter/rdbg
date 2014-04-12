package rdbg

import "time"

type DB struct{
	Name    Name
	Created time.Time
}

func (db *DB) Age() time.Duration {
	return time.Since(db.Created)
}

