package models

import (
	"database/sql"
	"fmt"
)

var createStudent = "INSERT INTO students (name, last_name) VALUES ($1, $2) returning name"
var getStudentByName = "SELECT name, last_name FROM students WHERE name=$1"
var updateStudentNameByName = "UPDATE students SET name=$1 WHERE name=$2"

type Student struct {
	Name     *string `json:"name" validate:"required"`
	LastName *string `json:"last_name"`
}

func (s *Student) CreateStudent(db *sql.DB) error {
	fmt.Println(s.Name)
	err := db.QueryRow(createStudent, *s.Name, *s.LastName).Scan(s.Name)

	if err != nil {
		return err
	}
	return nil
}

func (s *Student) GetStudentByName(db *sql.DB) error {
	fmt.Println(*s.Name)
	return db.QueryRow(getStudentByName, *s.Name).Scan(s.Name, s.LastName)
}

func (s *Student) UpdateStudentNameByName(db *sql.DB) error {
	_, err :=
		db.Exec(updateStudentNameByName, *s.Name)

	return err
}
