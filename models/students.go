package models

import "database/sql"

var createStudent = "INSERT INTO students (name, last_name) VALUES ($1, $2) returning name"
var getStudentByName = "SELECT name, last_name FROM students WHERE name=$1"
var updateStudentNameByName = "UPDATE students SET name=$1 WHERE name=$2"

type Student struct {
	Name     string `json:"name"`
	LastName string `json:"last_name" jsonschema:"required"`
}

func (s *Student) CreateStudent(db *sql.DB) error {
	err := db.QueryRow(createStudent, s.Name, s.LastName).Scan(&s.Name)

	if err != nil {
		return err
	}
	return nil
}

func (s *Student) GetStudentByName(db *sql.DB) error {
	return db.QueryRow(getStudentByName, s.Name).Scan(&s.Name, &s.LastName)
}

func (s *Student) UpdateStudentNameByName(db *sql.DB) error {
	_, err :=
		db.Exec(updateStudentNameByName, s.Name)

	return err
}
