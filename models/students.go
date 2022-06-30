package models

import (
	"database/sql"
	"errors"
	"fmt"
)

var createStudent = "INSERT INTO students (name, last_name) VALUES ($1, $2) returning name"
var getStudentByName = "SELECT name, last_name FROM students WHERE name=$1"
var updateStudentNameByName = "UPDATE students SET name=$1 WHERE name=$2"

var changeStudentPassword = `SELECT change_student_password(?, ?, ?) as shit`

// var changeStudentPassword = `SELECT change_student_password("9212001", "2744740129Me", "123456")`

type Student struct {
	StudentNO   string `json:"student_no" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

// func (s *Student) CreateStudent(db *sql.DB) error {
// 	fmt.Println(s.Name)
// 	err := db.QueryRow(createStudent, *s.Name, *s.LastName).Scan(s.Name)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *Student) GetStudentByName(db *sql.DB) error {
// 	fmt.Println(*s.Name)
// 	return db.QueryRow(getStudentByName, s.Name).Scan(&s.student_no, &s.LastName)
// }

// func (s *Student) UpdateStudentNameByName(db *sql.DB, newPassword string) error {
// 	res, err :=
// 		db.Exec(updateStudentNameByName, *s.student_no)

// 	fmt.Println()
// 	return err
// }

func (s *Student) ChangeStudentPassword(db *sql.DB) error {
	rowAffected := 0
	// err := db.QueryRow(changeStudentPassword).Scan(&rowAffected)
	err := db.QueryRow(changeStudentPassword, s.StudentNO, s.Password, s.NewPassword).Scan(&rowAffected)

	fmt.Println("function cal resssss", err, rowAffected)
	if rowAffected == 0 {
		return errors.New("No student found with provided credentials!")
	}

	return err
}
