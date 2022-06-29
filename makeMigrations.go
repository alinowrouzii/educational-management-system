package main

import (
	"database/sql"
	"log"
)

var createStudent = "CREATE TABLE students ( name VARCHAR(20), last_name VARCHAR(40) )"

func MakeMigrations(db *sql.DB) {
	execs := []struct {
		stmt       string
		shouldFail bool
	}{
		{stmt: createStudent},
	}
	for _, exec := range execs {
		_, err := db.Exec(exec.stmt)
		hasFailed := err != nil
		if exec.shouldFail != hasFailed {
			expected := "succeed"
			if exec.shouldFail {
				expected = "fail"
			}
			log.Printf("'%s' should have %sed but did not: %s", exec.stmt, expected, err)
		} else if exec.shouldFail {
			log.Printf("'%s' failed as expected: %s", exec.stmt, err)
		}
	}
	log.Println("finish!")
}
