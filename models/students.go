package models

import "database/sql"

var getStudentByName = "SELECT name, last_name FROM students WHERE id=$1"
var updateStudentNameByName = "UPDATE students SET name=$1 WHERE name=$2"

type Student struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func (s *Student) GetStudentByName(db *sql.DB) error {
	return db.QueryRow(getStudentByName, s.Name).Scan(&s.Name, &s.LastName)
}

func (s *Student) UpdateStudentNameByName(db *sql.DB) error {
	_, err :=
		db.Exec(updateStudentNameByName, s.Name)

	return err
}
